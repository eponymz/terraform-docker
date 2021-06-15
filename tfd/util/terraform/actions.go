package terraform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util/gitlab"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func Init(path string) int {
	var (
		fresh       = viper.GetBool("FRESH")
		initCommand = fmt.Sprintf("terraform init --upgrade=%t", fresh)
	)
	
	initExitCode := util.ExecExitCode(initCommand, path)
	return initExitCode
}

func Plan(path string, workspace string) int {
	var (
		planArgs        = []string{"-out=plan.tmp", "-detailed-exitcode"}
		planExit int    = 1
		tfVars   string = workspace
	)

	logrus.Tracef("Action: plan - called with args: %s", []string{path, workspace})
	if wsValid, varFile := WorkspaceExec(workspace, path); wsValid {
		if path != "." {
			tfVars = fmt.Sprintf("%s/%s", path, workspace)
		}
		if varFile {
			planArgs = append(planArgs, fmt.Sprintf("-var-file=%s.tfvars", tfVars))
		}
		planArgs = append(planArgs, path)
		logrus.Tracef("Terraform plan will execute against '%s' with: %s", path, planArgs)
		if planResult := util.ExecExitCode("terraform plan", planArgs...); planResult == 2 {
			if eval, returnString := PlanEval("plan.tmp"); eval {
				planExit = 0
			} else {
				logrus.Errorf("PlanEval failed: %s", returnString)
				planExit = 1
			}
		} else {
			planExit = planResult
		}
	}
	return planExit
}

func PlanEval(planFile string) (bool, string) {
	var (
		success               bool = true
		returnString          string
		destructiveChangePath string = `resource_changes.#(change.actions.#(=="delete"))#.address`
	)

	logrus.Infof("Changes detected in '%s' - Evaluating for destructive changes!", planFile)
	if _, err := os.Stat(planFile); os.IsNotExist(err) {
		returnString = fmt.Sprintf("Plan file '%s' does not exist!", planFile)
		success = false
	}

	planShow, err := exec.Command("terraform", "show", "-json", planFile).CombinedOutput()
	if err != nil {
		returnString = fmt.Sprintf("Terraform failed to output plan file '%s'!", planFile)
		success = false
	}

	if valid := gjson.Valid(string(planShow)); valid {
		destructiveChanges := gjson.Get(string(planShow), destructiveChangePath).Array()
		if len(destructiveChanges) > 0 {
			commentBody := fmt.Sprintf("%d destructive changes found! %s", len(destructiveChanges), destructiveChanges)
			logrus.Warn(commentBody)
			if commentErr := gitlab.PostMRComment(commentBody); commentErr != nil {
				returnString = fmt.Sprintf("PostMRComment failed: %s", commentErr)
			}
		} else {
			logrus.Infof("No destructive changes detected!")
		}
	} else {
		returnString = fmt.Sprintf("Plan file '%s' returned invalid JSON!", planFile)
		success = false
	}

	return success, returnString
}

func Apply(path string, workspace string) int {
	var (
		applyArgs   []string
		applyResult int = 1
	)

	isAutomation := viper.GetBool("AUTOAPPLY")
	logrus.Tracef("Running in CI/CD: %t", isAutomation)
	logrus.Tracef("Action: apply - called with args: %s", []string{path, workspace})

	if isAutomation {
		applyArgs = append(applyArgs, fmt.Sprintf("-auto-approve"))
	}

	if wsValid, varFile := WorkspaceExec(workspace, path); wsValid {
		if varFile {
			applyArgs = append(applyArgs, fmt.Sprintf("-var-file=%s.tfvars", workspace))
		}
		if path != "." {
			applyArgs = append(applyArgs, fmt.Sprintf("%s", path))
		}
		if varFile {
			logrus.Tracef("Terraform apply will execute with: %s", applyArgs)
			applyResult = util.ExecExitCode("terraform apply", applyArgs...)
		} else {
			logrus.Trace("Terraform apply will execute with no args. This is indicative of running in a default workspace.")
			applyResult = util.ExecExitCode("terraform apply", path)
		}
	}

	return applyResult
}

func WorkspaceSwitch(workspace string, path string) bool {
	var result bool = false
	if IsWorkspaceValid(workspace, path) {
		currentWorkspace, _ := exec.Command("terraform", "workspace", "show", path).CombinedOutput()
		if workspace != strings.Trim(string(currentWorkspace), "\n") {
			logrus.Infof("Switching to the %s workspace", workspace)
			switchWsExitCode := util.ExecExitCode("terraform workspace", "select", workspace, path)
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

func IsWorkspaceValid(workspace string, path string) bool {
	logrus.Tracef("Checking workspace '%s' validity.", workspace)
	var isValid bool = false
	wsList, _ := exec.Command("terraform", "workspace", "list", path).CombinedOutput()
	cleaned := util.CleanString(string(wsList))
	split := strings.Split(cleaned, " ")
	logrus.Trace(len(split))
	if len(split) > 1 && util.SliceContains(split, workspace) {
		isValid = true
	}
	logrus.Tracef("Available workspaces: %s, Requested workspace: %s, isValid returned: %t", split, workspace, isValid)
	return isValid
}

func WorkspaceExec(workspace string, path string) (bool, bool) {
	var (
		wsValid bool
		varFile bool
	)
	wd, _ := os.Getwd()
	if workspace == "default" {
		logrus.Warn("Default workspace has been selected! This can cause prompts for variables. If this occurs, check workspace selection!")
		logrus.Tracef("%s workspace provided. Skipping workspace validation for path: %s.", workspace, wd)
		wsValid = true
		varFile = false
	} else {
		wsValid = WorkspaceSwitch(workspace, path)
		varFile = true
	}
	return wsValid, varFile
}
