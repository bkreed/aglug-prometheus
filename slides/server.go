package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	http.Handle("/metrics", prometheus.Handler())                                    // HLone
	http.HandleFunc("/file", prometheus.InstrumentHandlerFunc("/file", fileHandler)) // HLtwo
	log.Fatal(http.ListenAndServe(":8080", nil))
}
