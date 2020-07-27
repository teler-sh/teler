package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/pkg/requests"
)

// ParseOptions will parse args/opts
func ParseOptions() *common.Options {
	options := &common.Options{}

	flag.StringVar(&options.ConfigFile, "f", "", "")
	flag.StringVar(&options.ConfigFile, "file", "", "")

	flag.IntVar(&options.Concurrency, "c", 20, "")
	flag.IntVar(&options.Concurrency, "concurrent", 20, "")

	flag.BoolVar(&options.Version, "v", false, "")
	flag.BoolVar(&options.Version, "version", false, "")

	// Override help flag
	flag.Usage = func() {
		h := []string{
			"",
			"Options:",
			"  -f, --file <FILE>           teler configuration file",
			"  -c, --concurrent <i>        Set the concurrency level to process log (default: 20)",
			"  -v, --version               Show current teler version",
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()

	// Show current version & exit
	if options.Version {
		showVersion()
	}

	// Check if stdin pipe was given
	options.Stdin = hasStdin()

	// Validates all given args/opts also for user teler config
	validate(options)

	// Getting all resources
	requests.Get(options)

	return options
}
