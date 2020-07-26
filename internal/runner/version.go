package runner

import (
	"fmt"
	"os"
)

func showVersion() {
	fmt.Printf("teler %s\n", version)
	os.Exit(2)
}
