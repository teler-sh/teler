package runner

import (
	"fmt"
	"os"
	"ktbs.dev/teler/versioninfo"
)

func showVersion() {
	fmt.Printf("teler %s\n", versioninfo.Version)
	os.Exit(2)
}
