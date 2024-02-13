package runner

import (
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/kitabisa/tailpipe"
	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/internal/alert"
	"github.com/kitabisa/teler/internal/event"
	"github.com/kitabisa/teler/pkg/errors"
	"github.com/kitabisa/teler/pkg/metrics"
	"github.com/kitabisa/teler/pkg/teler"
	"github.com/logrusorgru/aurora"
	"github.com/panjf2000/ants/v2"
	"github.com/projectdiscovery/gologger"
	"github.com/remeh/sizedwaitgroup"
	"github.com/satyrius/gonx"
)

// New read & pass stdin log
func New(options *common.Options) {
	var (
		reader *gonx.Reader
		input  *os.File
		pass   int
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
	sse := event.Run(options, common.Version)

	p, _ := ants.NewPoolWithFunc(con, func(line interface{}) {
		threat, obj := teler.Analyze(options, line.(*gonx.Entry))
		if threat {
			fmt.Printf("[%s] [%s] [%s] %s\n",
				aurora.Cyan(obj["time_local"]),
				aurora.Green(obj["remote_addr"]),
				aurora.Yellow(obj["category"]),
				aurora.Red(obj[obj["element"]]),
			)

			sse.Push(obj)
			log(options, obj)
			metrics.PrometheusInsert(options, obj)
			alert.New(options, common.Version, obj)
		}
		swg.Done()
	}, ants.WithPreAlloc(true))
	defer p.Release()

	go func() {
		for job := range jobs {
			swg.Add()
			_ = p.Invoke(job)
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

	if !options.Stdin && options.Follow {
		f, e := tailpipe.Open(options.Input)
		if e != nil {
			errors.Exit(e.Error())
		}

		reader = gonx.NewReader(f, format)
	} else {
		reader = gonx.NewReader(input, format)
	}

	for {
		line, err := reader.Read()
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
