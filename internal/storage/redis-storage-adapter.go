package storage

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_URI      = "REDIS_URI"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)

type RedisStorageAdapter struct {
	client *redis.Client
}

func NewRedisStorageAdapter() *RedisStorageAdapter {
	cl := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(REDIS_URI),
		Password: os.Getenv(REDIS_PASSWORD),
	})

	return &RedisStorageAdapter{
		client: cl,
	}
}

func (rsa RedisStorageAdapter) CountRequests(svcName string) int64 {
	ctx := context.Background()

	rq, err := rsa.client.LLen(ctx, svcName).Result()
	if err != nil {
		panic("Unable to retrieve requests from redis")
	}
	return rq
}

func (rsa RedisStorageAdapter) StoreRequest(svcName string) (string, error) {
	ctx := context.Background()
	t := time.Now()
	exp := t.Add(time.Second)
	status := rsa.client.Conn().Set(ctx, svcName, t.Unix(), time.Duration(exp.Unix()))
	return status.Result()
}
