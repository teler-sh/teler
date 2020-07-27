package runner

import (
	"os"
	"reflect"
	"strings"

	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/pkg/errors"
	"github.com/kitabisa/teler/pkg/matchers"
	"github.com/kitabisa/teler/pkg/parsers"
	"gopkg.in/validator.v2"
)

func validate(options *common.Options) {
	if !options.Stdin {
		errors.Exit("No stdin log to processed")
	}

	if options.ConfigFile == "" {
		errors.Exit("No config file specified")
	}

	config, errConfig := parsers.GetConfig(options.ConfigFile)
	if errConfig != nil {
		errors.Exit(errConfig.Error())
	}

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
