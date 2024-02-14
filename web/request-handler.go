package web

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/internal/balancer"
	"github.com/rs/zerolog/log"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("RequestHandler >> Handling request for path %s", r.URL.Path)
	blcr := balancer.NewBalancer(r.URL.Path)
	ds := blcr.Balance()
	ds.NewProxy().ServeHTTP(w, r)
}
