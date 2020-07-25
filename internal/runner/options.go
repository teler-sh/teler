package runner

import (
	"flag"

	e "github.com/kitabisa/teler/pkg/errors"
	"github.com/kitabisa/teler/pkg/parsers"
)

// Options contains the configuration options
type Options struct {
	Concurrency int    // Set the concurrent level
	ConfigFile  string // Specifies the config to use
	Stdin       bool   // Stdin specifies whether stdin input was given to the process
	Update      bool   // Updating resources to latest

	Config *parsers.Config
}

// ParseOptions will parse args/opts
func ParseOptions() *Options {
	options := &Options{}

	flag.StringVar(&options.ConfigFile, "f", "", "teler configuration file")
	flag.IntVar(&options.Concurrency, "c", 10, "Set the concurrency level")
	flag.BoolVar(&options.Update, "u", false, "Update teler resources to the latest")
	flag.Parse()

	// Check update flag
	if options.Update {
		// rsrc, _ := parsers.GetResources()
		// TODO
	}

	// Show user the banner
	showBanner()

	// Check if stdin pipe was given
	options.Stdin = hasStdin()

	val := validate(options)
	if val != nil {
		e.Err(val)
	}

	return options
}
