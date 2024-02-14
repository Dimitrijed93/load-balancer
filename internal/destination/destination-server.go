package destination

import (
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/rs/zerolog/log"
)

type Destination struct {
	Uri    *url.URL
	Mutext *sync.Mutex
}

func (d Destination) NewProxy() *httputil.ReverseProxy {
	log.Info().Msgf("New proxy for URI %s", d.Uri)
	return httputil.NewSingleHostReverseProxy(d.Uri)
}
