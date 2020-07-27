package matchers

import (
	"regexp"
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
