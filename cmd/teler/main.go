package main

import (
	"net/http"
	"runtime"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"ktbs.dev/teler/internal/runner"
	"ktbs.dev/teler/pkg/teler"
)

func init() {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu + 1)
}

func main() {
	// Parse the command line flags
	options := runner.ParseOptions()
	if options.Metrics {
		http.Handle("/metrics", promhttp.Handler())

		teler.MetricInit()
		go runner.New(options)

		http.ListenAndServe(":"+strconv.Itoa(options.MetricsPort), nil)

	} else {
		runner.New(options)
	}

}
