package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"go.uber.org/automaxprocs/maxprocs"
	"teler.app/internal/runner"
)

func init() {
	_, _ = maxprocs.Set()
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
}

func main() {
	// Parse the command line flags
	options := runner.ParseOptions()
	runner.New(options)
}
