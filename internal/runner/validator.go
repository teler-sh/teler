package runner

import (
	"os"
	"reflect"
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

func notification(options *common.Options) {
	config := options.Configs

	if config.Alert.Active {
		provider := strings.Title(config.Alert.Provider)
		field := reflect.ValueOf(&config.Notifications).Elem().FieldByName(provider)

		switch provider {
		case "Slack":
			field.FieldByName("URL").SetString(SlackAPI)
			matchers.IsHexcolor(field.FieldByName("Color").String())
			matchers.IsChannel(field.FieldByName("Channel").String())
		case "Telegram":
			field.FieldByName("URL").SetString(strings.Replace(TelegramAPI, ":token", field.FieldByName("Token").String(), -1))
			matchers.IsChatID(field.FieldByName("ChatID").String())
			matchers.IsParseMode(field.FieldByName("ParseMode").String())
		case "Discord":
			// TODO
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
