package errors

import (
	"os"
	"strings"

	log "github.com/projectdiscovery/gologger"
)

// Exit will display error details and stop the program
func Exit(err string) {
	if err != "" {
		for _, e := range strings.Split(strings.TrimSuffix(err, "\n"), "\n") {
			log.Errorf("Error! %s.\n", e)
		}
		os.Exit(1)
	}
}
