package roundrobin

import (
	"sync/atomic"

	destination "github.com/dimitrijed93/load-balancer/internal/destination"
	"github.com/rs/zerolog/log"
)

type RoundRobinBalancer struct {
	Destinations []destination.Destination
	next         uint32
}

func NewRoundRobinBalancer(d []destination.Destination) RoundRobinBalancer {
	return RoundRobinBalancer{
		Destinations: d,
		next:         0,
	}
}

func (rr RoundRobinBalancer) Balance() *destination.Destination {
	log.Info().Msgf("RoundRobinBalancer >> Balance using round robin, destinations: %s", rr.Destinations)
	n := atomic.AddUint32(&rr.next, 1)

	if int(n) > len(rr.Destinations) {
		atomic.StoreUint32(&rr.next, 0)
	}
	log.Info().Msgf("VALUE OF N %d", n)
	log.Info().Msgf("DD 1 %s", rr.Destinations[0])

	srv := rr.Destinations[(int(n)-1)%len(rr.Destinations)]
	return &srv
}
