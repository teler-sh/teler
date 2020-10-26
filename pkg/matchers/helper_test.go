package matchers

import "os"

func getTestEnvName(fnName string) string {
	return "TEST_" + fnName
}

func IsEnvSet(fnName string) bool {
	return os.Getenv(getTestEnvName(fnName)) == "1"
}

func SetTestEnv(fnName string) string {
	return "TEST_" + fnName + "=1"
}

func SetTestArgEnv(fnName string, arg string) string {
	return "TEST_ARG_" + fnName + "=" + arg
}

func GetTestArgEnv(fnName string) string {
	return os.Getenv("TEST_ARG_" + fnName)
}

func GetArgCmd(fnName string) string {
	return "-test.run=" + fnName
}
