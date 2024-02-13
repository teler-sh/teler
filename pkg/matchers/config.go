package matchers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kitabisa/teler/pkg/errors"
)

func errValidate(key string) {
	err := strings.Replace(errors.ErrConfigValidate, ":key", key, -1)
	errors.Exit(err)
}

// IsLogformat validates structured log format
func IsLogformat(s string) {
	if regexp := IsMatch(PatternLogformat, s); !regexp {
		errValidate("log format")
	}
}

// IsToken validates the token
func IsToken(s string) {
	if regexp := IsMatch(PatternToken, s); !regexp {
		errValidate("token")
	}
}

// IsColor validates the color code
func IsColor(s string) {
	if regexp := IsMatch(PatternColor, s); !regexp {
		errValidate("(hex) color")
	}
}

// IsChannel validates the channel for Slack & Discord
func IsChannel(s string) {
	if regexp := IsMatch(PatternChannel, s); !regexp {
		errValidate("channel")
	}
}

// IsWebhook validates the webhook URL for Slack & Discord
func IsWebhook(p string, s string) bool {
	var pat string

	switch p {
	case "Slack":
		pat = fmt.Sprintf(PatternWebhook, `hooks\.slack\.com`, `services\/.+`)
	case "Discord":
		pat = fmt.Sprintf(PatternWebhook, `discord\.com`, `api\/webhooks`)
	case "Mattermost":
		pat = fmt.Sprintf(PatternWebhook, `.+`, `hooks`)
	}

	return IsMatch(pat, s)
}

// IsChatID validates the chat_id for Telegram
func IsChatID(s string) {
	if _, isFloat := strconv.ParseFloat(s, 64); isFloat != nil {
		errValidate("chat_id")
	}
}

// IsCondition validates custom threat rules condition
func IsCondition(s string) {
	switch s {
	case "or", "and":
	default:
		errValidate("AND/OR for condition")
	}
}

// IsBlank validates nil field value
func IsBlank(s string, field string) {
	s = strings.TrimSpace(s)
	if s == "" {
		err := strings.Replace(errors.ErrBlankField, ":field", field, -1)
		errors.Exit(err)
	}
}
