package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	concurrency := 20
	jobs := make(chan string)
	var wg sync.WaitGroup

	flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
	flag.Parse()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for log := range jobs {
				fmt.Println(log)
				// Superman flying starts here
			}
			wg.Done()
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		jobs <- sc.Text()
	}

	close(jobs)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	wg.Wait()
}
