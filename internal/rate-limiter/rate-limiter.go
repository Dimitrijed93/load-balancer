package ratelimiter

import (
	"github.com/dimitrijed93/load-balancer/config"
	"github.com/dimitrijed93/load-balancer/internal/storage"
	"github.com/rs/zerolog/log"
)

type RateLimiter struct {
	config config.RateLimiterConfig
	sta    storage.StorageAdapter
}

func NewRateLimiter(rlc config.RateLimiterConfig) *RateLimiter {
	return &RateLimiter{
		config: rlc,
		sta:    storage.NewRedisStorageAdapter(),
	}
}

func (rl RateLimiter) RequestAllowed() bool {
	st, err := rl.sta.StoreRequest(rl.config.Name, rl.config.WindowSize)
	if err != nil {
		log.Error().Msgf("RateLimiter >> Storing request failed with %s \n", err)
	}
	log.Info().Msgf("RateLimiter >> Status of storing request %s \n", st)
	r := rl.sta.CountRequests(rl.config.Name)
	log.Info().Msgf("RateLimiter >> Request count for svc name %s is %d", rl.config.Name, r)
	return r < int(rl.config.MaxRequests)
}
