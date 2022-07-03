package runner

import (
	"fmt"
	"os"

	"teler.app/common"
)

func showVersion() {
	fmt.Printf("teler %s\n", common.Version)
	os.Exit(2)
}
