package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ktbs.dev/teler/common"
	"ktbs.dev/teler/pkg/cache"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/requests"
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

	flag.StringVar(&options.Output, "o", "", "")
	flag.StringVar(&options.Output, "output", "", "")

	flag.BoolVar(&options.Version, "v", false, "")
	flag.BoolVar(&options.Version, "version", false, "")

	flag.BoolVar(&options.RmCache, "rm-cache", false, "")

	flag.BoolVar(&options.JSON, "json", false, "")

	// Override help flag
	flag.Usage = func() {
		showBanner()
		h := []string{
			"",
			"Usage:",
			common.Usage,
			"",
			"Options:",
			"  -c, --config <FILE>         teler configuration file",
			"  -i, --input <FILE>          Analyze logs from data persistence rather than buffer stream",
			"  -x, --concurrent <i>        Set the concurrency level to analyze logs (default: 20)",
			"  -o, --output <FILE>         Save detected threats to file",
			"      --json                  Display threats in the terminal as JSON format",
			"      --rm-cache              Removes all cached resources",
			"  -v, --version               Show current teler version",
			"",
			"Examples:",
			common.Example,
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

	// Removes all cached resources on user-level directory
	if options.RmCache {
		rmCache()
	}

	// Check if stdin pipe was given
	options.Stdin = hasStdin()

	// Validates all given args/opts also for user teler config
	validate(options)

	// Check if resources is cached, then check
	// internet connection before get remote resources
	if !cache.Check() {
		if !isConnected() {
			errors.Exit("Check your internet connection")
		}
	}

	// Getting all resources
	requests.Resources(options)

	return options
}
