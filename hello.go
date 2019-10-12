package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	queryID = "id"
)

var (
	hiRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "custom_requests_hi_total",
		Help: "The total number of custom hi requests",
	}, []string{queryID})
)

func main() {
	// Runtime metrics automatically retrieved by prom client
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		ID := r.URL.Query().Get(queryID)

		hiRequests.With(prometheus.Labels{queryID: ID}).Inc()

		_, _ = fmt.Fprintf(w, "Hi %v\n", ID)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
