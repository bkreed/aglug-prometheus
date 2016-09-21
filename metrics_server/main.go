package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/file", prometheus.InstrumentHandlerFunc("/file", fileHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("files/%s", r.URL.Query().Get("name"))
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(404)
			w.Write([]byte("File not found\n"))
			return
		} else {
			w.WriteHeader(500)
			w.Write([]byte("Server Error\n"))
			return
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(buf)
}
