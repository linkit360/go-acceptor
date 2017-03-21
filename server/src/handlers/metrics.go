package handlers

import (
	"time"

	m "github.com/linkit360/go-utils/metrics"
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
