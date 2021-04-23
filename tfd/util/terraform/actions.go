package terraform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"tfd/util"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(path string) int {
	// wd, _ := os.Getwd()
	initExitCode := util.ExecExitCode("terraform init", path)
	return initExitCode
}

func Plan(path string, workspace string) int {
	var (
		planArgs string
		planExit int = 1
	)

	logrus.Tracef("Action: plan - called with args: %s", []string{path, workspace})
	planArgs = fmt.Sprint("-detailed-exitcode")

	if wsValid, addArgs := WorkspaceExec(workspace); wsValid {
		if addArgs != "" {
			planArgs += addArgs
		}
		logrus.Tracef("Terraform plan will execute with: %s", planArgs)
		if planResult := util.ExecExitCode("terraform plan", "-out=plan.tmp", planArgs, path); planResult == 2 {
			logrus.Debug("INF-155 will build out plan evaluation. Always returns 0 unless evaluation fails.")
			planExit = 0
		} else {
			planExit = planResult
		}
	}
	return planExit
}

func Apply(path string, workspace string) int {
	var (
		applyArgs   string
		applyResult int = 1
	)

	isAutomation := viper.GetBool("AUTOAPPLY")
	logrus.Tracef("Running in CI/CD: %t", isAutomation)
	logrus.Tracef("Action: apply - called with args: %s", []string{path, workspace})

	if isAutomation {
		applyArgs = fmt.Sprintf("-auto-approve")
	}

	if wsValid, addArgs := WorkspaceExec(workspace); wsValid {
		if addArgs != "" {
			applyArgs += addArgs
		}
		if applyArgs != "" {
			logrus.Tracef("Terraform apply will execute with: %s", applyArgs)
			applyResult = util.ExecExitCode("terraform apply", applyArgs, path)
		} else {
			logrus.Trace("Terraform apply will execute with no args. This is indicative of running in a default workspace.")
			applyResult = util.ExecExitCode("terraform apply", path)
		}
	}

	return applyResult
}

func WorkspaceSwitch(workspace string) bool {
	var result bool = false
	if IsWorkspaceValid(workspace) {
		currentWorkspace, _ := exec.Command("terraform", "workspace", "show").CombinedOutput()
		if workspace != strings.Trim(string(currentWorkspace), "\n") {
			logrus.Infof("Switching to the %s workspace", workspace)
			switchWsExitCode := util.ExecExitCode("terraform workspace", "select", workspace)
			if switchWsExitCode > 0 {
				logrus.Errorf("Failed to switch workspace to %s", workspace)
			} else {
				result = true
			}
		} else {
			logrus.Tracef("Already using the %s workspace", workspace)
			result = true
		}
	}
	return result
}

func IsWorkspaceValid(workspace string) bool {
	logrus.Tracef("Checking workspace '%s' validity.", workspace)
	var isValid bool = false
	wsList, _ := exec.Command("terraform", "workspace", "list").CombinedOutput()
	clean := strings.ReplaceAll(string(wsList), "* ", "")
	split := strings.Split(strings.TrimSpace(clean), "\n")
	if len(split) > 1 && util.SliceContains(split, workspace) {
		isValid = true
	}
	logrus.Tracef("Available workspaces: %s, Requested workspace: %s, isValid returned: %t", split, workspace, isValid)
	return isValid
}

func WorkspaceExec(workspace string) (bool, string) {
	var (
		wsValid bool
		addArgs string
	)
	wd, _ := os.Getwd()
	if workspace == "default" {
		logrus.Warn("Default workspace has been selected! This can cause prompts for variables. If this occurs, check workspace selection!")
		logrus.Tracef("%s workspace provided. Skipping workspace validation for path: %s.", workspace, wd)
		wsValid = true
	} else {
		wsValid = WorkspaceSwitch(workspace)
		addArgs = fmt.Sprintf(" -var-file=%s.tfvars", workspace)
	}
	return wsValid, addArgs
}
