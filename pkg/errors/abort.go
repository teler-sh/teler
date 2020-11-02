// +build !windows

package errors

import (
	"os"
	"syscall"
)

// Abort will terminate & sends SIGTERM to process
func Abort(i ...int) {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		os.Exit(i[0])
	}

	pgid, err := syscall.Getpgid(syscall.Getpid())
	if err != nil {
		Exit(err.Error())
	}

	// nolint:errcheck
	syscall.Kill(-pgid, syscall.SIGTERM)
}
