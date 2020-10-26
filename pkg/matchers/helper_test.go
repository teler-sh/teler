package matchers

import (
	"os"
	"os/exec"
)

func getTestEnvName(fnName string) string {
	return "TEST_" + fnName
}

func IsEnvSet(fnName string) bool {
	return os.Getenv(getTestEnvName(fnName)) == "1"
}

func GetTestArgEnv(fnName string) string {
	return os.Getenv("TEST_ARG_" + fnName)
}

func InitExecCommand(fnName string, text string) (err error) {
	c := NewCmd(fnName, text)
	cmd := exec.Command(os.Args[0], c.GetArgCmd())
	cmd.Env = append(os.Environ(), c.GetTestEnv(), c.GetTestArgEnv())
	err = cmd.Run()
	return
}

type Cmd struct {
	fnName string
	arg    string
}

func NewCmd(fn string, arg string) *Cmd {
	return &Cmd{fn, arg}
}

func (c Cmd) GetTestEnv() string {
	return "TEST_" + c.fnName + "=1"
}

func (c Cmd) GetTestArgEnv() string {
	return "TEST_ARG_" + c.fnName + "=" + c.arg
}

func (c Cmd) GetArgCmd() string {
	return "-test.run=" + c.fnName
}
