package tracing

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/util"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

type Tracing struct {
	handler http.Handler
}

func NewTracing(handler http.Handler) *Tracing {
	return &Tracing{handler: handler}
}

func (t Tracing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	correlationId := xid.New().String()
	log.Info().Msgf("Tracing >> CorrelationId %s", correlationId)
	w.Header().Add(util.HEADER_NAME_CORRELATION_ID, correlationId)
	t.handler.ServeHTTP(w, r)
}
