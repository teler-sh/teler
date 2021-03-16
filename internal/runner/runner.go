package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"regexp"

	"github.com/acarl005/stripansi"
	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/remeh/sizedwaitgroup"
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
	var input *os.File
	var out string
	var pass int

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
		gologger.Info().Msgf("Listening metrics on http://" + promserve + promendpoint)
	}

	jobs := make(chan *gonx.Entry)
	gologger.Info().Msg("Analyzing...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		<-stop
		gologger.Warning().Msg("Interuppted. Exiting...")

		close(jobs)
		done(pass)
	}()

	con := options.Concurrency
	swg := sizedwaitgroup.New(con)
	for i := 0; i < con; i++ {
		swg.Add()
		go func() {
			defer swg.Done()

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

					alert.New(options, common.Version, obj)
				}
			}
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
		pass++
	}

	close(jobs)
	swg.Wait()
	done(pass)
}

func done(i int) {
	if i == 0 {
		gologger.Warning().Msg("No logs analyzed, did you write log format correctly?")
	}
	gologger.Info().Msg("Done!")

	os.Exit(1)
}
