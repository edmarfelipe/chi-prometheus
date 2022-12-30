package chiprometheus

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func TestMiddleware(t *testing.T) {
	recorder := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Use(NewMiddleware("test"))

	r.Handle("/metrics", promhttp.Handler())
	r.Get(`/ok`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	r.Get(`/users/{firstName}`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	req1, err := http.NewRequest("GET", "http://localhost:3000/ok", nil)
	if err != nil {
		t.Error(err)
	}
	req2, err := http.NewRequest("GET", "http://localhost:3000/users/JoeBob", nil)
	if err != nil {
		t.Error(err)
	}
	req3, err := http.NewRequest("GET", "http://localhost:3000/users/Misty", nil)
	if err != nil {
		t.Error(err)
	}
	req4, err := http.NewRequest("GET", "http://localhost:3000/metrics", nil)
	if err != nil {
		t.Error(err)
	}

	r.ServeHTTP(recorder, req1)
	r.ServeHTTP(recorder, req2)
	r.ServeHTTP(recorder, req3)
	r.ServeHTTP(recorder, req4)

	body := recorder.Body.String()
	if !strings.Contains(body, reqsName) {
		t.Errorf("body does not contain request total entry '%s'", reqsName)
	}
	if !strings.Contains(body, latencyName) {
		t.Errorf("body does not contain request duration entry '%s'", latencyName)
	}

	req1Count := `chi_request_duration_milliseconds_count{code="OK",method="GET",path="/ok",service="test"} 1`
	req2Count := `chi_request_duration_milliseconds_count{code="OK",method="GET",path="/users/{firstName}",service="test"} 2`

	if !strings.Contains(body, req1Count) {
		t.Errorf("body does not contain req1 count summary '%s'", req1Count)
	}
	if !strings.Contains(body, req2Count) {
		t.Errorf("body does not contain req2 count summary '%s'", req2Count)
	}
}
