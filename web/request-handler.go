package web

import (
	"errors"
	"net/http"

	"github.com/dimitrijed93/load-balancer/internal/balancer"
	customerrors "github.com/dimitrijed93/load-balancer/internal/custom_errors"
	"github.com/rs/zerolog/log"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("RequestHandler >> Handling request for path %s", r.URL.Path)
	blcr, err := balancer.NewBalancer(r.URL.Path)

	var tooManyRequestError = &customerrors.RateLimitError{}
	log.Error().Msg("ERR " + tooManyRequestError.Error())

	if err != nil && errors.As(err, tooManyRequestError) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(err.Error()))
		return
	}
	ds := blcr.Balance()
	ds.NewProxy().ServeHTTP(w, r)
}
