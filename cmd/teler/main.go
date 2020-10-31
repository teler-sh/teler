package main

import (
	"net/http"
	"runtime"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"ktbs.dev/teler/internal/runner"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/metrics"
)

func init() {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu + 1)
}

func main() {
	// Parse the command line flags
	options := runner.ParseOptions()

	//run metric
	http.Handle("/metrics", promhttp.Handler())
	metrics.MetricInit()
	go runner.New(options)

	err := http.ListenAndServe(":"+strconv.Itoa(options.Metrics), nil)
	if err != nil {
		errors.Exit(err.Error())
	}
}
