package ratelimiter

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/config"
	"github.com/dimitrijed93/load-balancer/internal/storage"
	"github.com/rs/zerolog/log"
)

type RateLimiter struct {
	sta     storage.StorageAdapter
	handler http.Handler
}

func NewRateLimiter(handler http.Handler) *RateLimiter {
	return &RateLimiter{
		sta:     storage.NewRedisStorageAdapter(),
		handler: handler,
	}
}

func (rl RateLimiter) requestAllowed(config *config.RateLimiterConfig) bool {
	st, err := rl.sta.StoreRequest(config.Name, config.WindowSize)
	if err != nil {
		log.Error().Msgf("RateLimiter >> Storing request failed with %s \n", err)
	}
	log.Info().Msgf("RateLimiter >> Status of storing request %s \n", st)
	r := rl.sta.CountRequests(config.Name)
	log.Info().Msgf("RateLimiter >> Request count for svc name %s is %d", config.Name, r)
	return r <= int(config.MaxRequests)
}

func (rl RateLimiter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config := config.NewRateLimiterConfig(r.URL.Path)
	allow := rl.requestAllowed(config)
	log.Info().Msgf("RateLimiter >> Request is allowed %s", allow)
	if !allow {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("RateLimiter >> Too many requests"))
		return
	}
	rl.handler.ServeHTTP(w, r)
}
