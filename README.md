# go-metrics-example
Go Observability example with Prometheus and Grafana

This is a simple implementation of an observability layer using Prometheus and Grafana implemented in go.

### Prometheus
Prometheus is the metric collector and it offers few metric options that can be used in specific situations. See https://prometheus.io/docs/concepts/metric_types/ for more details about the metric types.

Prometheus works as a pull-based service that will hit /metrics endpoint to retrieve the metrics addressed. We can configure how often it does that under the config yaml file > scrape_interval property.
It also automatically generates Go metrics such as number of goroutines, memory usage etc. All of this through the prometheus go client.

### Grafana
Grafana is our dashboard where we integrate with Prometheus and build graphs with that data. With that we can generate useful dashboards that helps identifying outages, annomalies, rps, latency and so on.

<img width="1504" alt="Screenshot 2023-05-24 at 11 07 49 AM" src="https://github.com/felipeaz/go-observability-example/assets/32846823/98c50b09-79ec-44fa-ae17-b85387708f85">
