package alert

import (
	"reflect"

	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/matchers"
	"ktbs.dev/teler/pkg/utils"
)

// New will initialize notification provider & send threat alerts
func New(options *common.Options, version string, log map[string]string) {
	var token string
	var webhooked bool

	config := options.Configs
	if config.Alert.Active {
		provider := utils.Title(config.Alert.Provider)
		field := reflect.ValueOf(&config.Notifications).Elem().FieldByName(provider) // nosemgrep

		if matchers.IsWebhook(provider, field.FieldByName("Webhook").String()) {
			token = field.FieldByName("Webhook").String()
			webhooked = true
		} else {
			token = field.FieldByName("Token").String()
		}

		switch provider {
		case "Slack":
			toSlack(
				token,
				field.FieldByName("Channel").String(),
				field.FieldByName("Color").String(),
				log,
				webhooked,
			)
		case "Telegram":
			toTelegram(
				token,
				field.FieldByName("ChatID").String(),
				log,
			)
		case "Discord":
			toDiscord(
				token,
				field.FieldByName("Channel").String(),
				field.FieldByName("Color").String(),
				version,
				log,
				webhooked,
			)
		case "Mattermost":
			toMattermost(
				token,
				field.FieldByName("Color").String(),
				log,
			)
		}
	}
}
