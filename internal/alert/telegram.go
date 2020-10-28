package alert

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func toTelegram(token string, chatId string, parseMode string, log map[string]string) {
	id, err := strconv.Atoi(chatId)
	if err != nil {
		panic(err)
	}

	api, err := telegramBot.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	message := telegramBot.NewMessage(int64(id), telegramMessage(log))
	message.ParseMode = parseMode

	time.Sleep(1 * time.Second)
	_, err = api.Send(message)
	if err != nil {
		time.Sleep(5 * time.Second)
	}
}

func telegramMessage(log map[string]string) string {
	var buffer bytes.Buffer

	template, err := template.ParseFiles(fmt.Sprintf("internal/alert/template/telegram.tmpl"))
	if err != nil {
		panic(err)
	}

	err = template.Execute(&buffer, log)
	if err != nil {
		panic(err)
	}

	return buffer.String()
}
