package matchers

import (
	"errors"
	"regexp"
	"strings"

	e "github.com/kitabisa/teler/pkg/errors"
)

func regExp(pattern string, s string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func errValidate(key string) error {
	err := strings.Replace(e.ErrConfigValidate, ":key", key, -1)
	return errors.New(err)
}

// IsToken validates the token
func IsToken(s string) error {
	if regexp := regExp(PatternToken, s); !regexp {
		return errValidate("token")
	}

	return nil
}

// IsHexcolor validates the hex color code
func IsHexcolor(s string) error {
	if regexp := regExp(PatternHexcolor, s); !regexp {
		return errValidate("hex color")
	}

	return nil
}
