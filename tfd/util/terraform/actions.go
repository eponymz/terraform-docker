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
	os.Chdir(path)
	wd, _ := os.Getwd()
	initExitCode := util.ExecExitCode("terraform init", wd)
	return initExitCode
}

func Plan(path string, workspace string) int {
	var (
		planArgs string
		planExit int
	)

	logrus.Tracef("Action: plan - called with args: %s", []string{path, workspace})
	os.Chdir(path)
	planArgs = fmt.Sprint("-detailed-exitcode")

	wsValid, addArgs := ExecWorkspaceValidate(workspace)
	if addArgs != "" {
		planArgs += addArgs
	}

	if !wsValid {
		return 1
	}

	logrus.Tracef("Terraform plan will execute with: %s", planArgs)
	planResult := util.ExecExitCode("terraform plan", "-out=plan.tmp", planArgs)
	switch planResult {
	case 0:
	case 1:
		planExit = planResult
		break
	case 2:
		// this case will eventually call a function to evaluate changes and notify of destructive actions
		// will ALWAYS return 0 unless function fails for visibility of plan changes
		planExit = 0
		break
	}
	return planExit
}

func Apply(path string, workspace string) int {
	var (
		applyArgs   string
		applyResult int
	)

	isAutomation := viper.GetBool("AUTOAPPLY")
	logrus.Tracef("Running in CI/CD: %t", isAutomation)
	logrus.Tracef("Action: apply - called with args: %s", []string{path, workspace})

	os.Chdir(path)

	if isAutomation {
		applyArgs = fmt.Sprintf("-auto-approve")
	}

	wsValid, addArgs := ExecWorkspaceValidate(workspace)
	if addArgs != "" {
		applyArgs += addArgs
	}

	if !wsValid {
		return 1
	}

	if applyArgs != "" {
		logrus.Tracef("Terraform apply will execute with: %s", applyArgs)
		applyResult = util.ExecExitCode("terraform apply", applyArgs)
	} else {
		logrus.Trace("Terraform apply will execute with no args. This is indicative of running in a default workspace.")
		applyResult = util.ExecExitCode("terraform apply")
	}

	return applyResult
}

// Switches to valid workspace
func Workspace(path string, workspace string) bool {
	var result bool
	os.Chdir(path)
	if WorkspaceCheck(path, workspace) {
		currentWorkspace, _ := exec.Command("terraform", "workspace", "show").CombinedOutput()
		if workspace != strings.Trim(string(currentWorkspace), "\n") {
			logrus.Infof("Switching to the %s workspace", workspace)
			switchWsExitCode := util.ExecExitCode("terraform workspace", "select", workspace)
			if switchWsExitCode > 0 {
				logrus.Errorf("Failed to switch workspace to %s", workspace)
				result = false
			} else {
				result = true
			}
		} else {
			logrus.Tracef("Already using the %s workspace", workspace)
		}
	}
	return result
}

// checks if workspace provided is a valid option to switch to
func WorkspaceCheck(path string, workspace string) bool {
	logrus.Tracef("Checking workspace '%s' validity.", workspace)
	var isValid bool = false
	wsList, _ := exec.Command("terraform", "workspace", "list").CombinedOutput()
	clean := strings.ReplaceAll(string(wsList), "* ", "")
	split := strings.Split(strings.TrimSpace(clean), "\n")
	if len(split) > 1 {
		if !util.SliceContains(split, workspace) {
			logrus.Errorf("%s is not a valid workspace! Available workspaces are: %s", workspace, clean)
		} else {
			logrus.Tracef("%s is a valid workspace!", workspace)
			isValid = true
		}
	} else {
		logrus.Errorf("Path: %s does not utilize custom workspaces! Did you mean to pass 'default' as the workspace?", path)
	}
	return isValid
}

// executes the workspace validation and switch flow
func ExecWorkspaceValidate(workspace string) (bool, string) {
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
		wsValid = Workspace(wd, workspace)
		addArgs = fmt.Sprintf(" -var-file=%s.tfvars", workspace)
	}
	return wsValid, addArgs
}
