package web

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/config"
	"github.com/dimitrijed93/load-balancer/internal/balancer"
	"github.com/rs/zerolog/log"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("RequestHandler >> Handling request for path %s", r.URL.Path)

	cnf := config.NewServiceConfig(r.URL.Path)

	blcr, err := balancer.NewBalancer(r.URL.Path, cnf)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	ds := blcr.Balance()
	ds.NewProxy().ServeHTTP(w, r)
}
