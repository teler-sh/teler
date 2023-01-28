package common

import (
	"os"

	"teler.app/pkg/parsers"
)

// Options contains the configuration options
type Options struct {
	Concurrency int              // Set the concurrent level
	ConfigFile  string           // Specifies the config to use
	Configs     *parsers.Configs // Get teler configuration interface
	Follow      bool             // Specify if the logs should be streamed
	Input       string           // Parse log from data persistence rather than buffer stream
	JSON        bool             // Display threats in the terminal as JSON format
	Output      *os.File         // Write log output into file
	RmCache     bool             // To remove all cached resources on local
	Stdin       bool             // Stdin specifies whether stdin input was given to the process
	Version     bool             // Version check of teler flag
}
