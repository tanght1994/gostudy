package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type LabelValues []string

func (lvs LabelValues) With(labelValues ...string) LabelValues {
	if len(labelValues)%2 != 0 {
		labelValues = append(labelValues, "unknown")
	}
	return append(lvs, labelValues...)
}

func main() {
	fieldKeys := []string{"fieldKeys1", "fieldKeys2"}

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "tht_namespace",
		Subsystem: "tht_subsystem",
		Name:      "tht_counter_name",
		Help:      "tht_help",
	}, fieldKeys)
	prometheus.MustRegister(counter)

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "tht_namespace",
		Subsystem: "tht_subsystem",
		Name:      "tht_histogram_name",
		Help:      "tht_help",
	}, fieldKeys)
	prometheus.MustRegister(histogram)

	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "tht_namespace",
		Subsystem: "tht_subsystem",
		Name:      "tht_summary_name",
		Help:      "tht_help",
	}, fieldKeys)
	prometheus.MustRegister(summary)

	lables := prometheus.Labels{}
	lables["fieldKeys1"] = "value1"
	lables["fieldKeys2"] = "value2"

	counter.With(lables).Add(1)
	counter.With(lables).Add(1)
	counter.With(lables).Add(1)

	histogram.With(lables).Observe(1)
	histogram.With(lables).Observe(1)
	histogram.With(lables).Observe(5)
	histogram.With(lables).Observe(1)
	histogram.With(lables).Observe(3)
	histogram.With(lables).Observe(3)
	histogram.With(lables).Observe(1)
	histogram.With(lables).Observe(1)
	histogram.With(lables).Observe(6)
	histogram.With(lables).Observe(1)

	summary.With(lables).Observe(1)
	summary.With(lables).Observe(2)
	summary.With(lables).Observe(3)
	summary.With(lables).Observe(4)
	summary.With(lables).Observe(5)
	summary.With(lables).Observe(6)
	summary.With(lables).Observe(7)
	summary.With(lables).Observe(8)
	summary.With(lables).Observe(9)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
