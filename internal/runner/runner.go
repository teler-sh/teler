package runner

import (
	"fmt"
	"io"
	"os"
	"sync"

	gonx "github.com/satyrius/gonx"
)

func New(options *Options) {
	jobs := make(chan *gonx.Entry)
	var wg sync.WaitGroup

	fmt.Println()

	for i := 0; i < options.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for log := range jobs {
				fmt.Println(log)
				// Superman flying starts here
			}
			wg.Done()
		}()
	}

	config := options.Config.Configs
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
