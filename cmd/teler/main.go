package main

import (
	"runtime"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"ktbs.dev/teler/internal/runner"
)

func init() {
	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu + 1)
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
}

func main() {
	// Parse the command line flags
	options := runner.ParseOptions()
	runner.New(options)
}
