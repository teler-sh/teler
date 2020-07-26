package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"

	e "github.com/kitabisa/teler/pkg/errors"
	"github.com/kitabisa/teler/pkg/parsers"
)

// Options contains the configuration options
type Options struct {
	Concurrency int    // Set the concurrent level
	ConfigFile  string // Specifies the config to use
	Stdin       bool   // Stdin specifies whether stdin input was given to the process
	Update      bool   // Updating resources to latest
	Version     bool   // Updating resources to latest

	Config *parsers.Config
}

// ParseOptions will parse args/opts
func ParseOptions() *Options {
	options := &Options{}

	flag.StringVar(&options.ConfigFile, "f", "", "")
	flag.StringVar(&options.ConfigFile, "file", "", "")

	flag.IntVar(&options.Concurrency, "c", 20, "")
	flag.IntVar(&options.Concurrency, "concurrent", 20, "")

	flag.BoolVar(&options.Update, "u", false, "")
	flag.BoolVar(&options.Update, "update", false, "")

	flag.BoolVar(&options.Version, "v", false, "")
	flag.BoolVar(&options.Version, "version", false, "")

	// Override help flag
	flag.Usage = func() {
		h := []string{
			"",
			"Options:",
			"  -f, --file <FILE>           teler configuration file",
			"  -c, --concurrent <i>        Set the concurrency level to process log (default: 20)",
			"  -u, --update                Update all teler resources to the latest",
			"  -v, --version               Show current teler version",
			"",
		}

		showBanner()
		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()

	// Check update flag
	if options.Update {
		// rsrc, _ := parsers.GetResources()
		// TODO
	}

	// Show current version & exit
	if options.Version {
		showVersion()
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
