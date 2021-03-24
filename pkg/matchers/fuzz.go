package matchers

import "github.com/sahilm/fuzzy"

func IsMatchFuzz(pattern string, s []string) bool {
	matches := fuzzy.Find(pattern, s)

	return len(matches) > 0
}
