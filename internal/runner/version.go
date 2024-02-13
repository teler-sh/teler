package runner

import (
	"fmt"
	"os"

	"github.com/kitabisa/teler/common"
)

func showVersion() {
	fmt.Printf("teler %s\n", common.Version)
	os.Exit(2)
}
