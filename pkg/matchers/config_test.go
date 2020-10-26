package matchers

import (
	"os"
	"os/exec"
	"testing"
)

func TestIsLogformat(t *testing.T) {
	fnName := "TestIsLogformat"
	if IsEnvSet(fnName) {
		IsLogformat(GetTestArgEnv(fnName))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestIsLogformat")
	cmd.Env = append(os.Environ(), "TEST_TestIsLogformat=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
