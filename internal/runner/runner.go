package runner

import (
	"io"
	"os"
	"regexp"
	"sync"

	log "github.com/projectdiscovery/gologger"
	"github.com/satyrius/gonx"
	"ktbs.me/teler/common"
	"ktbs.me/teler/pkg/errors"
	"ktbs.me/teler/pkg/teler"
)

func init() {
	if !isConnected() {
		errors.Exit("Check your internet connection")
	}
}

func removeLBR(s string) string {
	re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000A}\x{000B}\x{000C}\x{000D}\x{0085}\x{2028}\x{2029}]`)
	return re.ReplaceAllString(s, ``)
}

// New read & pass stdin log
func New(options *common.Options) {
	var wg sync.WaitGroup
	var input *os.File
	jobs := make(chan *gonx.Entry)
	log.Infof("Analyzing...")

	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for log := range jobs {
				teler.Analyze(options, log)
			}
			wg.Done()
		}()
	}

	if options.Stdin {
		input = os.Stdin
	} else {
		f, e := os.Open(options.Input)
		if e != nil {
			errors.Exit(e.Error())
		}
		input = f
	}

	config := options.Configs
	format := removeLBR(config.Logformat)
	buffer := gonx.NewReader(input, format)
	for {
		line, err := buffer.Read()
		if err == io.EOF {
			break
		}
		jobs <- line
	}

	close(jobs)

	wg.Wait()
	log.Infof("Done!")
}
