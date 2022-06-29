package runner

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/projectdiscovery/gologger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/metrics"
)

func metric(options *common.Options) {
	m := options.Configs.Metrics
	v := reflect.ValueOf(m)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).FieldByName("Active").Bool() {
			switch t.Field(i).Name {
			case "Prometheus":
				startPrometheus(options)
			}
		}
	}
}

func startPrometheus(options *common.Options) {
	p := options.Configs.Metrics.Prometheus

	if p.Host == "" {
		p.Host = "127.0.0.1"
	}

	if p.Port == 0 {
		p.Port = 9090
	}

	if p.Endpoint == "" {
		p.Endpoint = "/metrics"
	}

	s := fmt.Sprint(p.Host, ":", strconv.Itoa(p.Port))
	e := p.Endpoint

	go func() {
		http.Handle(e, promhttp.Handler())

		err := http.ListenAndServe(s, nil) // nosemgrep
		if err != nil {
			errors.Exit(err.Error())
		}
	}()

	metrics.Prometheus()
	gologger.Info().Msgf(fmt.Sprint("Listening metrics on http://", s, e))
}
