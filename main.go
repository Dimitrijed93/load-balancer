package main

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/web"
	"github.com/dimitrijed93/load-balancer/web/debug"
)

func main() {
	http.HandleFunc("/debug", debug.RedisDebugHandler)
	http.HandleFunc("/", web.RequestHandler)

	// http.Handle("/", mux)
	http.ListenAndServe(":3030", nil)
}
