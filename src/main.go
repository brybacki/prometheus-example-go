package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestCounter.WithLabelValues(r.URL.Path).Inc()
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func namedHander(name string) http.HandlerFunc {
	requestCounter.WithLabelValues(name).Inc()

	return updateHandler
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	val, _ := strconv.ParseFloat(r.URL.Path, 64)
	gauge.WithLabelValues("user").Set(val)
	fmt.Fprintf(w, "Hello %s!", r.URL.Path)
}

var (
	promRegistry   = prometheus.NewRegistry() // new clean registry w/o go metrics
	requestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "example_go_request_count",
			Help: "Requests total",
		},
		[]string{"path"})
	gauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "example_go_gauge",
			Help: "Arbitraty gauge",
		},
		[]string{"updated"})
)

func main() {
	promRegistry.MustRegister(requestCounter)
	promRegistry.MustRegister(gauge)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{}))
	mux.HandleFunc("/", handler)
	mux.Handle(
		"/update/",
		http.StripPrefix("/update/", namedHander("/update/")))
	http.ListenAndServe(":8123", mux)
	fmt.Println("Helo")
}
