package runner

import "net/http"

func isConnected() bool {
	_, err := http.Get(Google204)
	if err != nil {
		return false
	}
	return true
}
