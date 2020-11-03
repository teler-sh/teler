package common

import (
	"os"

	"ktbs.dev/teler/pkg/parsers"
)

// Options contains the configuration options
type Options struct {
	Concurrency int              // Set the concurrent level
	ConfigFile  string           // Specifies the config to use
	Stdin       bool             // Stdin specifies whether stdin input was given to the process
	Version     bool             // Version check of teler flag
	Input       string           // Parse log from data persistence rather than buffer stream
	Output      string           // Save detected threats to file
	OutFile     *os.File         // Write log output into file
	Configs     *parsers.Configs // Get teler configuration interface
	JSON        bool             // Display threats in the terminal as JSON format
	RmCache     bool             // To remove all cached resources on local
}
