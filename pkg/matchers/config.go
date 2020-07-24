package matchers

import (
	"errors"
	"regexp"
)

func regExp(pattern string, s string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}

// IsToken validates the token
func IsToken(s string) error {
	regexp := regExp(`^(xox[p|b|o|a]-\d{10,12}-\d{12}-\w+)|(\d{9}:[a-zA-Z0-9_-]{35})$`, s)
	if !regexp {
		return errors.New("Only validates token; please check your config file")
	}

	return nil
}

// IsHexcolor validates the hex color code
func IsHexcolor(s string) error {
	regexp := regExp(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`, s)
	if !regexp {
		return errors.New("Only validates hex color; please check your config file")
	}

	return nil
}
