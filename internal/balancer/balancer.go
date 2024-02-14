package balancer

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	customerrors "github.com/dimitrijed93/load-balancer/internal/custom_errors"
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

func NewBalancer(uri string) (Balancer, error) {
	cp, err := config.NewConfigParser(os.Getenv(util.ENV_VAR_CONFIG_PATH))
	cnf := cp.Config
	if err != nil {
		log.Error().Msg("Unable to open config")
	}

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
		log.Info().Msgf("Balancer >> RateLimiterConfig not empty. WindowSize = %d MaxRequests = %d ServiceName = %s",
			svc.RateLimiterConfig.WindowSize, svc.RateLimiterConfig.MaxRequests, svc.RateLimiterConfig.Name)
		rl := ratelimiter.NewRateLimiter(svc.RateLimiterConfig)
		if !rl.RequestAllowed() {
			return nil, customerrors.RateLimitError{Limit: int(svc.RateLimiterConfig.MaxRequests)}
		}
	}

	switch svc.Type {
	case util.LOAD_BALANCER_TYPE_ROUND_ROBIN:
		return roundrobin.NewRoundRobinBalancer(dst), nil
	default:
		return roundrobin.NewRoundRobinBalancer(dst), nil
	}
}
