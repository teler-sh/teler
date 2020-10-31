// +build windows

package errors

import "os"

// Abort specifies the os.Exit function
var Abort = os.Exit
