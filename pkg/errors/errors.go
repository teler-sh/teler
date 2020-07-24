package errors

import (
	"os"
	"strings"

	log "github.com/projectdiscovery/gologger"
)

// Err will display error details and stop the program
func Err(err error) {
	if err != nil {
		s := err.Error()
		for _, e := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
			log.Errorf("Error! %s.\n", e)
		}
		os.Exit(1)
	}
}
