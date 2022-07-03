package alert

import (
	"bytes"
	"html/template"
	"strconv"

	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api"
	"teler.app/pkg/errors"
)

func toTelegram(token string, chatID string, log map[string]string) {
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		errors.Show(err.Error())
	}

	api, err := telegramBot.NewBotAPI(token)
	if err != nil {
		errors.Exit(err.Error())
	}

	message := telegramBot.NewMessage(id, telegramMessage(log))
	message.ParseMode = "MarkdownV2"

	// TODO: Displays an error if it does not exceed the rate-limit
	// nolint:errcheck
	api.Send(message)
}

func telegramMessage(log map[string]string) string {
	var buffer bytes.Buffer

	tpl, err := template.ParseFiles("internal/alert/template/telegram.tmpl")
	if err != nil {
		errors.Exit(err.Error())
	}

	err = tpl.Execute(&buffer, log)
	if err != nil {
		errors.Exit(err.Error())
	}

	return buffer.String()
}
