package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/acarl005/stripansi"
	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/remeh/sizedwaitgroup"
	"github.com/satyrius/gonx"
	"ktbs.dev/teler/common"
	"ktbs.dev/teler/internal/alert"
	"ktbs.dev/teler/pkg/errors"
	"ktbs.dev/teler/pkg/metrics"
	"ktbs.dev/teler/pkg/teler"
)

// New read & pass stdin log
func New(options *common.Options) {
	var (
		input *os.File
		out   string
		pass  int
	)

	go metric(options)

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
	go func() {
		for log := range jobs {
			swg.Add()
			go func(line *gonx.Entry) {
				defer swg.Done()

				threat, obj := teler.Analyze(options, line)
				if threat {
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
					metrics.Send(obj)
				}
			}(log)
		}
	}()

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
