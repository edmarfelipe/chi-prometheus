package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/edmarfelipe/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := chi.NewRouter()

	// if you want to use other buckets than the default (300, 1200, 5000) you can run:
	// chiprometheus.NewMiddleware("serviceName", 400, 1600, 700)
	r.Use(chiprometheus.NewMiddleware("test_service"))
	r.Handle("/metrics", promhttp.Handler())

	r.Get("/ok", func(w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "slept %d milliseconds\n", sleep)
	})

	r.Get("/users/{firstName}", func(w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "slept %d milliseconds\n", sleep)
	})

	r.Get("/users/{id}/contacts", func(w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "slept %d milliseconds\n", sleep)
	})

	r.Get("/notfound", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "not found")
	})

	http.ListenAndServe(":3000", r)
}
