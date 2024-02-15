package storage

import (
	"context"
	"os"
	"time"

	"github.com/dimitrijed93/load-balancer/util"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	MAX_NUMBER_OF_REQUESTS = 1000
)

type RedisStorageAdapter struct {
	client *redis.Client
}

func NewRedisStorageAdapter() *RedisStorageAdapter {
	redisUri := os.Getenv(util.ENV_VAR_REDIS_URI)
	redisPwd := os.Getenv(util.ENV_VAR_REDIS_PASSWORD)

	log.Info().Msgf("RedisStorageAdapter >> Redis URI %s", redisUri)

	cl := redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: redisPwd,
	})

	return &RedisStorageAdapter{
		client: cl,
	}
}

func (rsa RedisStorageAdapter) CountRequests(svcName string) int {
	log.Info().Msgf("RedisStorageAdapter >> Counting request for svc %s", svcName)
	ctx := context.Background()
	var cursor uint64

	r, _, err := rsa.client.Conn().Scan(ctx, cursor, svcName+"*", MAX_NUMBER_OF_REQUESTS).Result()

	if err != nil {
		panic(err)
	}
	c := len(r)
	log.Info().Msgf("RedisStorageAdapter >> Count of request for svc %s is %d", svcName, c)
	return c
}

func (rsa RedisStorageAdapter) StoreRequest(svcName string, exp int) (string, error) {
	ctx := context.Background()
	log.Info().Msgf("RedisStorageAdapter >> Storing request for service name %s", svcName)

	now := time.Now().Unix()
	d := time.Duration(exp * int(time.Millisecond))
	log.Info().Msgf("EXPIRE %s ", d.String())
	r, err := rsa.client.Conn().Set(ctx, svcName+time.Now().String(),
		now, d).Result()

	if err != nil {
		panic(err.Error())
	}

	return r, err
}
