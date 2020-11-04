package runner

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/validator.v2"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/matchers"
	"ktbs.dev/teler/pkg/parsers"
)

func validate(options *common.Options) {
	if !options.Stdin {
		if options.Input == "" {
			errors.Exit(errors.ErrNoInputLog)
		}
	}

	if options.ConfigFile == "" {
		telerEnv := os.Getenv("TELER_CONFIG")
		if telerEnv == "" {
			errors.Exit(errors.ErrNoInputConfig)
		} else {
			options.ConfigFile = telerEnv
		}
	}

	if options.Output != "" {
		f, errOutput := os.OpenFile(options.Output,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if errOutput != nil {
			errors.Exit(errOutput.Error())
		}
		options.OutFile = f
	}

	config, errConfig := parsers.GetConfig(options.ConfigFile)
	if errConfig != nil {
		errors.Exit(errConfig.Error())
	}

	// Validates log format
	matchers.IsLogformat(config.Logformat)
	options.Configs = config

	// Validates notification parts on configuration files
	notification(options)

	if errVal := validator.Validate(options); errVal != nil {
		errors.Exit(errVal.Error())
	}
}

func prometheus(options *common.Options) (bool, string, string) {
	config := options.Configs
	if config.Prometheus.Active {
		if config.Prometheus.Host == "" {
			config.Prometheus.Host = "127.0.0.1"
		}

		if config.Prometheus.Port == 0 {
			config.Prometheus.Port = 9090
		}

		if config.Prometheus.Endpoint == "" {
			config.Prometheus.Endpoint = "/metrics"
		}
	}

	server := config.Prometheus.Host + ":" + strconv.Itoa(config.Prometheus.Port)

	return config.Prometheus.Active, server, config.Prometheus.Endpoint
}

func notification(options *common.Options) {
	config := options.Configs

	if config.Alert.Active {
		provider := strings.Title(config.Alert.Provider)
		field := reflect.ValueOf(&config.Notifications).Elem().FieldByName(provider)

		switch provider {
		case "Slack", "Discord":
			matchers.IsColor(field.FieldByName("Color").String())
			matchers.IsChannel(field.FieldByName("Channel").String())
		case "Telegram":
			matchers.IsChatID(field.FieldByName("ChatID").String())
		default:
			errors.Exit(strings.Replace(errors.ErrAlertProvider, ":platform", config.Alert.Provider, -1))
		}

		matchers.IsToken(field.FieldByName("Token").String())
	}
}

func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}
