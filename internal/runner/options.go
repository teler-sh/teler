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

	flag.StringVar(&options.ConfigFile, "c", "", "")
	flag.StringVar(&options.ConfigFile, "config", "", "")

	flag.StringVar(&options.Input, "i", "", "")
	flag.StringVar(&options.Input, "input", "", "")

	flag.IntVar(&options.Concurrency, "x", 20, "")
	flag.IntVar(&options.Concurrency, "concurrent", 20, "")

	flag.BoolVar(&options.Version, "v", false, "")
	flag.BoolVar(&options.Version, "version", false, "")

	// Override help flag
	flag.Usage = func() {
		showBanner()
		h := []string{
			"",
			"Usage:",
			usage,
			"",
			"Options:",
			"  -c, --config <FILE>         teler configuration file",
			"  -i, --input <FILE>          Analyze logs from data persistence rather than buffer stream",
			"  -x, --concurrent <i>        Set the concurrency level to analyze logs (default: 20)",
			"  -v, --version               Show current teler version",
			"",
			"Examples:",
			example,
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()

	// Show current version & exit
	if options.Version {
		showVersion()
	}

	// Show the banner to user
	showBanner()

	// Check if stdin pipe was given
	options.Stdin = hasStdin()

	// Validates all given args/opts also for user teler config
	validate(options)

	// Getting all resources
	requests.Resources(options)

	return options
}
