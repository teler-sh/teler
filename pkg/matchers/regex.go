package matchers

import "regexp"

func IsMatch(pattern string, s string) bool {
	defer func() {
		recover()
	}()

	re := regexp.MustCompile(pattern)
	if re.FindString(s) != "" {
		return true
	}

	return false
}
