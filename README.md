# chi-prometheus

[Prometheus](http://prometheus.io) middleware for [chi v5](https://go-chi.io/). 

## Usage

```go
func main() {
    r := chi.NewRouter()
    
    r.Use(chiprometheus.NewMiddleware("service_name"))
    r.Handle("/metrics", promhttp.Handler())
	
   // others routers ...
	
    http.ListenAndServe(":3000", r)
}
```


## What do you get

An endpoint with the following information:

```sh
# HELP chi_pattern_request_duration_milliseconds How long it took to process the request, partitioned by status code, method and HTTP path (with patterns).
# TYPE chi_pattern_request_duration_milliseconds histogram
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="test_service",le="300"} 1
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="test_service",le="1200"} 1
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="test_service",le="5000"} 1
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/metrics",service="test_service",le="+Inf"} 1
chi_pattern_request_duration_milliseconds_sum{code="OK",method="GET",path="/metrics",service="test_service"} 0.975926
chi_pattern_request_duration_milliseconds_count{code="OK",method="GET",path="/metrics",service="test_service"} 1
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/users/{firstName}",service="test_service",le="300"} 0
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/users/{firstName}",service="test_service",le="1200"} 0
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/users/{firstName}",service="test_service",le="5000"} 2
chi_pattern_request_duration_milliseconds_bucket{code="OK",method="GET",path="/users/{firstName}",service="test_service",le="+Inf"} 2
chi_pattern_request_duration_milliseconds_sum{code="OK",method="GET",path="/users/{firstName}",service="test_service"} 4755.052755
chi_pattern_request_duration_milliseconds_count{code="OK",method="GET",path="/users/{firstName}",service="test_service"} 2
# HELP chi_pattern_requests_total How many HTTP requests processed, partitioned by status code, method and HTTP path (with patterns).
# TYPE chi_pattern_requests_total counter
chi_pattern_requests_total{code="OK",method="GET",path="/metrics",service="test_service"} 1
chi_pattern_requests_total{code="OK",method="GET",path="/users/{firstName}",service="test_service"} 2
```