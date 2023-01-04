package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	sleeptime := int64(10)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		res := int64(0)
		for {
			val := atomic.LoadInt64(&sleeptime)
			if val >= 20 {
				res = val
				break
			}
			newval := val + 1
			if atomic.CompareAndSwapInt64(&sleeptime, val, newval) {
				res = newval
				break
			}
		}
		w.Write([]byte(fmt.Sprintf("%d", res)))
	})
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		res := int64(0)
		for {
			val := atomic.LoadInt64(&sleeptime)
			if val <= 0 {
				res = val
				break
			}
			newval := val - 1
			if atomic.CompareAndSwapInt64(&sleeptime, val, newval) {
				res = newval
				break
			}
		}
		w.Write([]byte(fmt.Sprintf("%d", res)))
	})
	go http.ListenAndServe(":8080", nil)

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "tht_counter_name",
		Help: "tht_help",
	}, []string{"k1", "k2"})
	prometheus.MustRegister(counter)

	// histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
	// 	Name:    "tht_histogram_name",
	// 	Help:    "tht_help",
	// 	Buckets: []float64{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
	// }, []string{"k1", "k2"})
	// prometheus.MustRegister(histogram)

	// summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
	// 	Name:       "tht_summary_name",
	// 	Help:       "tht_help",
	// 	Objectives: map[float64]float64{0.5: 0, 0.6: 0, 0.7: 0, 0.8: 0, 0.9: 0, 1: 0},
	// }, []string{"k1", "k2"})
	// prometheus.MustRegister(summary)

	lables1 := prometheus.Labels{"k1": "1", "k2": "1"}
	lables2 := prometheus.Labels{"k1": "2", "k2": "2"}

	for {
		counter.With(lables1).Add(1)
		// val := rand.Float64() * 100
		// histogram.With(lables1).Observe(val)
		// summary.With(lables1).Observe(val)

		if rand.Float64() > 0.5 {
			counter.With(lables2).Add(1)
		} else {
			counter.With(lables2).Add(2)
		}

		tmp := atomic.LoadInt64(&sleeptime) * 10
		sleep := rand.Intn(int(10+tmp)) + 10
		sleep = sleep * int(time.Millisecond)
		time.Sleep((time.Duration(sleep)))
	}
}
