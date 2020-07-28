package runner

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/kitabisa/teler/common"
	"github.com/kitabisa/teler/pkg/errors"
	"github.com/satyrius/gonx"
)

func init() {
	if !isConnected() {
		errors.Exit("Check your internet connection")
	}
}

// New read & pass stdin log
func New(options *common.Options) {
	var wg sync.WaitGroup
	var input *os.File

	jobs := make(chan *gonx.Entry)

	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for log := range jobs {
				fmt.Printf("%+v", log)
				// Superman flying starts here
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
	format := config.Logformat
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
}
