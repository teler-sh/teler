package matchers

import "regexp"

func IsMatch(pattern string, s string) bool {
	defer func() {
		_ = recover()
	}()

	re := regexp.MustCompile(pattern)
	return re.FindString(s) != ""
}
