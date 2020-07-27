package errors

import (
	"os"
	"strings"

	log "github.com/projectdiscovery/gologger"
)

// Exit will display error details and stop the program
func Exit(err error) {
	if err != nil {
		s := err.Error()
		for _, e := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
			log.Errorf("Error! %s.\n", e)
		}
		os.Exit(1)
	}
}
