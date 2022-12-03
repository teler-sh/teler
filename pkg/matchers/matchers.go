package matchers

import (
	"regexp"
	"strings"

	"github.com/sahilm/fuzzy"
)

func IsAny(substr string, s string) bool {
	if strings.Index(s, substr) > 0 {
		return false
	}

	return true
}

func IsMatch(pattern string, s string) bool {
	defer func() {
		_ = recover()
	}()

	re := regexp.MustCompile(pattern)
	return re.FindString(s) != ""
}

func IsMatchFuzz(pattern string, s []string) bool {
	matches := fuzzy.Find(pattern, s)

	return len(matches) > 0
}
