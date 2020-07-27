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
	showBanner()
	if !isConnected() {
		errors.Exit("Check your internet connection")
	}
}

// New read & pass stdin log
func New(options *common.Options) {
	jobs := make(chan *gonx.Entry)
	var wg sync.WaitGroup

	fmt.Println()

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

	config := options.Configs
	format := config.Format
	buffer := gonx.NewReader(os.Stdin, format)
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
