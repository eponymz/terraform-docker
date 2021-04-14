package terraform

import (
	"os"
	"tfd/util"
)

func Init(path string) int {
	os.Chdir(path)
	wd, _ := os.Getwd()
	initExitCode := util.ExecExitCode("terraform init", wd)
	return initExitCode
}
