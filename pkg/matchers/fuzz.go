package matchers

import "github.com/sahilm/fuzzy"

func IsMatchFuzz(pattern string, s string) bool {
	matches := fuzzy.Find(pattern, s)
	if len(matches) > 0 {
		return true
	}

	return false
}
