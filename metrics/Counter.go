package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)


var addrCounter = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
var (
	taskCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "RH",
		Subsystem: "demo",
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func init() {
	prometheus.MustRegister(taskCounter)
}


func main() {
	flag.Parse()
	go func() {
		for {
			taskCounter.Inc()
			time.Sleep(time.Second * 1)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addrCounter, nil))
}
