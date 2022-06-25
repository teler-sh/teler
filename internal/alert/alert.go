package alert

import (
	"reflect"

	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/utils"
)

// New will initialize notification provider & send threat alerts
func New(options *common.Options, version string, log map[string]string) {
	config := options.Configs

	if config.Alert.Active {
		provider := utils.Title(config.Alert.Provider)
		field := reflect.ValueOf(&config.Notifications).Elem().FieldByName(provider)

		switch provider {
		case "Slack":
			toSlack(
				field.FieldByName("Token").String(),
				field.FieldByName("Channel").String(),
				field.FieldByName("Color").String(),
				log,
			)
		case "Telegram":
			toTelegram(
				field.FieldByName("Token").String(),
				field.FieldByName("ChatID").String(),
				log,
			)
		case "Discord":
			toDiscord(
				field.FieldByName("Token").String(),
				field.FieldByName("Channel").String(),
				field.FieldByName("Color").String(),
				version,
				log,
			)
		}
	}
}
