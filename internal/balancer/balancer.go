package balancer

import (
	"fmt"
	"net/url"

	destination "github.com/dimitrijed93/load-balancer/internal/destination"

	"github.com/dimitrijed93/load-balancer/config"
	roundrobin "github.com/dimitrijed93/load-balancer/internal/balancer/round-robin"
	"github.com/dimitrijed93/load-balancer/util"
	"github.com/rs/zerolog/log"
)

type Balancer interface {
	Balance() *destination.Destination
}

func NewBalancer(uri string, cnf config.ServiceConfig) (Balancer, error) {
	// cp, err := config.NewConfigParser(os.Getenv(util.ENV_VAR_CONFIG_PATH))
	// cnf := cp.Config
	// if err != nil {
	// 	log.Error().Msg("Unable to open config")
	// }

	// var svc config.ServiceConfig
	// for _, item := range cnf.ServicesConfig {
	// 	if strings.HasPrefix(uri, item.Path) {
	// 		svc = item
	// 	}
	// }

	var dst []destination.Destination

	for _, item := range cnf.ServersConfig {
		u := fmt.Sprintf("%s:%d%s", item.Host, item.Port, uri)
		log.Info().Msgf("Balancer >> Build destination URI %s", u)
		var targetUri *url.URL
		targetUri, _ = url.Parse(u)
		d := destination.Destination{Uri: targetUri}
		dst = append(dst, d)
	}

	log.Info().Msgf("Balancer >> New balancer of type %s", cnf.Type)

	// if !svc.RateLimiterConfig.IsEmpty() {
	// 	log.Info().Msgf("Balancer >> RateLimiterConfig not empty. WindowSize = %d MaxRequests = %d ServiceName = %s",
	// 		svc.RateLimiterConfig.WindowSize, svc.RateLimiterConfig.MaxRequests, svc.RateLimiterConfig.Name)
	// 	rl := ratelimiter.NewRateLimiter(svc.RateLimiterConfig)
	// 	if !rl.RequestAllowed() {
	// 		return nil, customerrors.RateLimitError{Limit: int(svc.RateLimiterConfig.MaxRequests)}
	// 	}
	// }

	switch cnf.Type {
	case util.LOAD_BALANCER_TYPE_ROUND_ROBIN:
		return roundrobin.NewRoundRobinBalancer(dst), nil
	default:
		return roundrobin.NewRoundRobinBalancer(dst), nil
	}
}
