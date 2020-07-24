package runner

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/kitabisa/teler/pkg/matchers"
	"github.com/kitabisa/teler/pkg/parsers"
	"gopkg.in/validator.v2"
)

func validate(options *Options) error {
	if !options.Stdin {
		return errors.New("No stdin log to processed")
	}

	if options.ConfigFile == "" {
		return errors.New("No config file specified")
	}

	config, errConf := parsers.GetConfig(options.ConfigFile)
	if errConf != nil {
		return errConf
	}

	options.Config = config

	if notif := options.notification(); notif != nil {
		return notif
	}

	if errVal := validator.Validate(options); errVal != nil {
		return errVal
	}

	return nil
}

func (options *Options) notification() error {
	config := options.Config.Configs
	notif := reflect.ValueOf(&options.Config.Notifications)

	if config.Notification.Active {
		provider := strings.Title(config.Notification.Provider)
		field := notif.Elem().FieldByName(provider)

		switch provider {
		case "Slack":
			field.FieldByName("URL").SetString(SlackAPI)
		case "Telegram":
			field.FieldByName("URL").SetString(strings.Replace(TelegramAPI, ":token", field.FieldByName("Token").String(), -1))
			if field.FieldByName("ChatID").String() == "" {
				return errors.New("Telegram \"chat_id\" is not set")
			}
		default:
			return errors.New("Provider \"" + config.Notification.Provider + "\" not available")
		}

		if errToken := notToken(field.FieldByName("Token").String()); errToken != nil {
			return errToken
		}
	}

	return nil
}

func notToken(s string) error {
	return matchers.IsToken(s)
}

func notHexcolor(s string) error {
	return matchers.IsHexcolor(s)
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
