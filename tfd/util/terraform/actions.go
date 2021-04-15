package terraform

import (
	"os"
	"os/exec"
	"strings"
	"tfd/util"

	"github.com/sirupsen/logrus"
)

func Init(path string) int {
	os.Chdir(path)
	wd, _ := os.Getwd()
	initExitCode := util.ExecExitCode("terraform init", wd)
	return initExitCode
}

// func Plan() {}

// func Apply() {}

// Used to validate the workspace value provided and switch to it.
func Workspace(path string, workspace string) int {
	var wsReturnCode int = 0 // used to determine moving forward or not in downstream functions
	os.Chdir(path)
	wsList, _ := exec.Command("terraform", "workspace", "list").CombinedOutput()
	clean := strings.ReplaceAll(string(wsList), "* ", "")
	split := strings.Split(clean, "\n")
	if !util.SliceContains(split, workspace) {
		logrus.Errorf("%s is not a valid workspace! Available workspaces are: %s", workspace, clean)
		wsReturnCode = 1
	} else {
		logrus.Tracef("%s is a valid workspace!", workspace)
	}
	currentWorkspace, _ := exec.Command("terraform", "workspace", "show").CombinedOutput()
	if workspace != strings.Trim(string(currentWorkspace), "\n") {
		logrus.Infof("Switching to the %s workspace", workspace)
		switchWsExitCode := util.ExecExitCode("terraform workspace", "select", workspace)
		if switchWsExitCode > 0 {
			logrus.Errorf("Failed to switch workspace to %s", workspace)
			wsReturnCode = 1
		}
	} else {
		logrus.Tracef("Already using the %s workspace", workspace)
	}
	return wsReturnCode
}
