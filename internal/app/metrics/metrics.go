package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Properties struct {
	ID          string
	Name        string
	Namespace   string
	Description string
	Type        string
	Label       string
	Properties  []string
}

func NewCounter(p Properties) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      p.Name,
			Namespace: p.Namespace,
			Help:      p.Description,
		},
		[]string{p.Label},
	)
}

func NewHistogram(p Properties) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      p.Name,
			Namespace: p.Namespace,
			Help:      p.Description,
			Buckets:   prometheus.LinearBuckets(0.01, 0.05, 10),
		},
		[]string{p.Label},
	)
}
