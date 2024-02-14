package main

import (
	"net/http"

	"github.com/dimitrijed93/load-balancer/web"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", web.RequestHandler)

	http.Handle("/", mux)
	http.ListenAndServe(":3030", nil)
}
