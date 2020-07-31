package matchers

import "regexp"

func IsMatch(pattern string, s string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}
