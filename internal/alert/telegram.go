package alert

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/projectdiscovery/gologger"
)

func toTelegram(token string, chatID string, parseMode string, log map[string]string) {
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		gologger.Errorf(err.Error())
	}

	api, err := telegramBot.NewBotAPI(token)
	if err != nil {
		gologger.Errorf(err.Error())
	}

	message := telegramBot.NewMessage(id, telegramMessage(log))
	message.ParseMode = parseMode

	for {
		_, err = api.Send(message)
		time.Sleep(1 * time.Second)
		if err == nil {
			break
		}
	}
}

func telegramMessage(log map[string]string) string {
	var buffer bytes.Buffer

	template, err := template.ParseFiles(fmt.Sprintf("internal/alert/template/telegram.tmpl"))
	if err != nil {
		gologger.Warningf(err.Error())
	}

	err = template.Execute(&buffer, log)
	if err != nil {
		gologger.Warningf(err.Error())
	}

	return buffer.String()
}
