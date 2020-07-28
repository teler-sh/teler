package matchers

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kitabisa/teler/pkg/errors"
)

func regExp(pattern string, s string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func errValidate(key string) {
	err := strings.Replace(errors.ErrConfigValidate, ":key", key, -1)
	errors.Exit(err)
}

// IsLogformat validates structured log format
func IsLogformat(s string) {
	if regexp := regExp(PatternLogformat, s); !regexp {
		errValidate("log format")
	}
}

// IsToken validates the token
func IsToken(s string) {
	if regexp := regExp(PatternToken, s); !regexp {
		errValidate("token")
	}
}

// IsHexcolor validates the hex color code
func IsHexcolor(s string) {
	if regexp := regExp(PatternHexcolor, s); !regexp {
		errValidate("hex color")
	}
}

// IsParseMode validates the parse mode for Telegram
func IsParseMode(s string) {
	if regexp := regExp(PatternParseMode, s); !regexp {
		errValidate("parse mode [Mardown(v2), or HTML]")
	}
}

// IsChannel validates the channel for Slack
func IsChannel(s string) {
	if regexp := regExp(PatternChannel, s); !regexp {
		errValidate("Slack channel")
	}
}

// IsChatID validates the chat_id for Slack
func IsChatID(s string) {
	if _, isFloat := strconv.ParseFloat(s, 8); isFloat != nil {
		errValidate("chat_id")
	}
}
