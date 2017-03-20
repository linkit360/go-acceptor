package handlers

import (
	m "github.com/linkit360/go-utils/metrics"
	"time"
)

var (
	success m.Gauge
	errors  m.Gauge
)

func InitMetrics() {
	success = m.NewGauge("", "", "success", "success")
	errors = m.NewGauge("", "", "errors", "errors")

	go func() {
		for range time.Tick(time.Minute) {
			success.Update()
			errors.Update()
		}
	}()
}
