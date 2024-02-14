package ratelimiter

import (
	"github.com/dimitrijed93/load-balancer/config"
	"github.com/dimitrijed93/load-balancer/internal/storage"
	"github.com/rs/zerolog/log"
)

type RateLimiter struct {
	config  config.RateLimiterConfig
	sta     storage.StorageAdapter
	svcName string
}

func NewRateLimiter(rlc config.RateLimiterConfig, svcName string) *RateLimiter {
	return &RateLimiter{
		config:  rlc,
		sta:     storage.RedisStorageAdapter{},
		svcName: svcName,
	}
}

func (rl RateLimiter) Limit() {
	rl.sta.StoreRequest(rl.svcName)
	r := rl.sta.CountRequests(rl.svcName)
	log.Info().Msgf("RateLimiter >> Request count for svc name %s is %d", rl.svcName, r)
}
