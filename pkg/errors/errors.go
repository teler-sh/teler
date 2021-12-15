package errors

import (
	"bufio"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// Exit will show error details and stop the program
func Exit(err string) {
	if err != "" {
		count := 0
		lines := bufio.NewScanner(strings.NewReader(err))

		for lines.Scan() {
			var msg string

			if count == 0 {
				msg = "Error! "
			}
			msg += strings.TrimSpace(lines.Text())

			Show(msg)
			count++
		}

		gologger.Info().Msgf("Use \"-h\" flag for more info about command.")
		Abort(9)
	}
}

// Show error message
func Show(msg string) {
	gologger.Error().Msg(msg)
}
