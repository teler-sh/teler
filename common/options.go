package common

import "github.com/kitabisa/teler/pkg/parsers"

// Options contains the configuration options
type Options struct {
	Concurrency int              // Set the concurrent level
	ConfigFile  string           // Specifies the config to use
	Stdin       bool             // Stdin specifies whether stdin input was given to the process
	Version     bool             // Version check of teler flag
	Input       string           // Parse log from data persistence rather than buffer stream
	Configs     *parsers.Configs // Get teler configuration interface
}
