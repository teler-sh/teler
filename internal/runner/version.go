package runner

import (
	"fmt"
	"os"

	"ktbs.dev/teler/common"
)

func showVersion() {
	fmt.Printf("teler %s\n", common.Version)
	os.Exit(2)
}
