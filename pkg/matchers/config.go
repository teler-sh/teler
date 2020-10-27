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

// IsHexcolor validates the hex color code
func IsHexcolor(s string) {
	if regexp := IsMatch(PatternHexcolor, s); !regexp {
		errValidate("hex color")
	}
}

// IsParseMode validates the parse mode for Telegram
func IsParseMode(s string) {
	if regexp := IsMatch(PatternParseMode, s); !regexp {
		errValidate("parse mode [Markdown(v2), or HTML]")
	}
}

// IsChannel validates the channel for Slack
func IsChannel(s string) {
	if regexp := IsMatch(PatternChannel, s); !regexp {
		errValidate("Slack channel")
	}
}

// IsChatID validates the chat_id for Telegram
func IsChatID(s string) {
	if _, isFloat := strconv.ParseFloat(s, 64); isFloat != nil {
		errValidate("chat_id")
	}
}
