package roundrobin

import (
	"sync/atomic"

	destination "github.com/dimitrijed93/load-balancer/internal/destination"
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
	n := atomic.AddUint32(&rr.next, 1)

	if int(n) > len(rr.Destinations) {
		atomic.StoreUint32(&rr.next, 0)
	}
	srv := rr.Destinations[(int(n)-1)%len(rr.Destinations)]
	return &srv
}
