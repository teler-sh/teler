package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/acarl005/stripansi"
	"github.com/logrusorgru/aurora"
	log "github.com/projectdiscovery/gologger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/satyrius/gonx"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/internal/alert"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/metrics"
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
	var out string

	metric, promserve, promendpoint := prometheus(options)
	if metric {
		go func() {
			http.Handle(promendpoint, promhttp.Handler())

			err := http.ListenAndServe(promserve, nil)
			if err != nil {
				errors.Exit(err.Error())
			}
		}()

		metrics.Init()
		log.Infof("Listening metrics on http://" + promserve + promendpoint)
	}

	jobs := make(chan *gonx.Entry)
	log.Infof("Analyzing...")

	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for log := range jobs {
				threat, obj := teler.Analyze(options, log)

				if threat {
					if metric {
						metrics.GetThreatTotal.WithLabelValues(obj["category"]).Inc()
					}

					if options.JSON {
						json, err := json.Marshal(obj)
						if err != nil {
							errors.Exit(err.Error())
						}
						out = fmt.Sprintf("%s\n", json)
					} else {
						out = fmt.Sprintf("[%s] [%s] [%s] %s\n",
							aurora.Cyan(obj["time_local"]),
							aurora.Green(obj["remote_addr"]),
							aurora.Yellow(obj["category"]),
							aurora.Red(obj[obj["element"]]),
						)
					}

					fmt.Print(out)

					if options.Output != "" {
						if !options.JSON {
							out = stripansi.Strip(out)
						}

						if _, write := options.OutFile.WriteString(out); write != nil {
							errors.Show(write.Error())
						}
					}

					alert.New(options, version, obj)
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
