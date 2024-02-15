package main

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/config"
	ratelimiter "github.com/dimitrijed93/load-balancer/internal/middleware/rate-limiter"
	"github.com/dimitrijed93/load-balancer/web"
	"github.com/dimitrijed93/load-balancer/web/debug"
)

func main() {
	config.NewLoadBalancerConfig()

	mux := http.NewServeMux()

	mux.HandleFunc("/debug", debug.RedisDebugHandler)
	mux.HandleFunc("/", web.RequestHandler)

	rlm := ratelimiter.NewRateLimiter(mux)

	http.ListenAndServe(":3030", rlm)
}
