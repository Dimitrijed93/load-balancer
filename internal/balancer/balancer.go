package balancer

import (
	"fmt"
	"net/url"
	"strings"

	destination "github.com/dimitrijed93/load-balancer/internal/destination"
	ratelimiter "github.com/dimitrijed93/load-balancer/internal/rate-limiter"

	"github.com/dimitrijed93/load-balancer/config"
	roundrobin "github.com/dimitrijed93/load-balancer/internal/balancer/round-robin"
	"github.com/dimitrijed93/load-balancer/util"
	"github.com/rs/zerolog/log"
)

type Balancer interface {
	Balance() *destination.Destination
}

func NewBalancer(uri string) Balancer {
	cp := config.FileSystemConfigParser{}
	cnf, err := cp.Parse("examples/example.yaml")
	if err != nil {
		log.Error().Msg("Unable to open config")
	}

	log.Info().Msgf("Balancer >> Found configuration %s", cnf.ServicesConfig)

	var svc config.ServiceConfig
	for _, item := range cnf.ServicesConfig {
		if strings.HasPrefix(uri, item.Path) {
			svc = item
		}
	}

	if svc.Path == util.EMPTY_STRING {
		panic("Invalid path, no configuration found for path " + uri)
	}

	var dst []destination.Destination

	for _, item := range svc.ServersConfig {
		u := fmt.Sprintf("%s:%d%s", item.Host, item.Port, uri)
		log.Info().Msgf("Balancer >> Build destination URI %s", u)
		var targetUri *url.URL
		targetUri, err = url.Parse(u)
		d := destination.Destination{Uri: targetUri}
		dst = append(dst, d)
	}

	log.Info().Msgf("Balancer >> New balancer of type %s", svc.Type)

	if !svc.RateLimiterConfig.IsEmpty() {
		rl := ratelimiter.NewRateLimiter(svc.RateLimiterConfig, svc.Name)
		rl.Limit()
	}

	switch svc.Type {
	case util.LOAD_BALANCER_TYPE_ROUND_ROBIN:
		return roundrobin.NewRoundRobinBalancer(dst)
	default:
		return roundrobin.NewRoundRobinBalancer(dst)
	}
}
