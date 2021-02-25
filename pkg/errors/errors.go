package errors

import (
	"strings"

	"github.com/projectdiscovery/gologger"
)

// Exit will show error details and stop the program
func Exit(err string) {
	msg := "Error! "
	if err != "" {
		for _, e := range strings.Split(strings.TrimSuffix(err, "\n"), "\n") {
			msg += e
			Show(msg)
		}
		gologger.Info().Msgf("Use \"-h\" flag for more info about command.")
		Abort(9)
	}
}

// Show error message
func Show(msg string) {
	gologger.Error().Msgf("%s\n", msg)
}
