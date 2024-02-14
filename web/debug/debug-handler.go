package debug

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/internal/storage"
	"github.com/rs/zerolog/log"
)

const (
	REQUEST_PARAM_SERVICE_NAME = "svc"
)

func RedisDebugHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get(REQUEST_PARAM_SERVICE_NAME)
	rsa := storage.NewRedisStorageAdapter()
	log.Info().Msgf("DebugHandler >> Handling request for svc %s", p)
	redisData := rsa.CountRequests(p)

	log.Info().Msgf("DebugHandler >> redis count %d", redisData)

	// blcr := balancer.NewBalancer(r.URL.Path)
	// ds := blcr.Balance()
	// ds.NewProxy().ServeHTTP(w, r)
}
