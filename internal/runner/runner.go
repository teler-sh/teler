package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"sync"

	"github.com/logrusorgru/aurora"
	log "github.com/projectdiscovery/gologger"
	"github.com/satyrius/gonx"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/internal/alert"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/teler"
)

func removeLBR(s string) string {
	re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000A}\x{000B}\x{000C}\x{000D}\x{0085}\x{2028}\x{2029}]`)
	return re.ReplaceAllString(s, ``)
}

// New read & pass stdin log
func New(options *common.Options) {
	var wg sync.WaitGroup
	var input *os.File

	if options.Metrics {
		log.Infof("Listening metrics on port : " + strconv.Itoa(options.MetricsPort))
	}

	jobs := make(chan *gonx.Entry)
	log.Infof("Analyzing...")

	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for log := range jobs {
				threat, obj := teler.Analyze(options, log)

				if threat {
					fmt.Printf("[%s] [%s] [%s] %s\n",
						aurora.Cyan(obj["time_local"]),
						aurora.Green(obj["remote_addr"]),
						aurora.Yellow(obj["category"]),
						aurora.Red(obj[obj["element"]]),
					)

					if options.Output != "" {
						json, _ := json.Marshal(obj)
						_, write := options.OutFile.WriteString(
							fmt.Sprintf("%s\n", json),
						)
						if write != nil {
							errors.Show(write.Error())
						}
					}

					alert.New(options, obj)
				}
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
