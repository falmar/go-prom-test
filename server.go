package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_amount",
		Help: "Number of requests",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(reqCount)
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		reqCount.Inc()
	})

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
