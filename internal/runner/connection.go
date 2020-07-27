package runner

import "net/http"

func isConnected() bool {
	if _, err := http.Get(Google204); err != nil {
		return false
	}
	return true
}
