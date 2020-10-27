package errors

import (
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// ExitFunc specifies the os.Exit function
var ExitFunc = os.Exit

// Exit will show error details and stop the program
func Exit(err string) {
	msg := "Error! "
	if err != "" {
		for _, e := range strings.Split(strings.TrimSuffix(err, "\n"), "\n") {
			msg += e
			Show(msg)
		}
		gologger.Infof("Use \"-h\" flag for more info about command.")
		ExitFunc(1)
	}
}

// Show error message
func Show(msg string) {
	gologger.Errorf("%s\n", msg)
}
