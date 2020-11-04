package matchers

import (
	"strconv"
	"strings"

	"ktbs.dev/teler/pkg/errors"
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

// IsChatID validates the chat_id for Telegram
func IsChatID(s string) {
	if _, isFloat := strconv.ParseFloat(s, 64); isFloat != nil {
		errValidate("chat_id")
	}
}
