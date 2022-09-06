package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/cmd/deploy"
	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
)

func TestDefaultAction(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-p=deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "No changes."
	if !strings.Contains(got, wants) {
		// t.Fatalf("tf.Plan() should run as default when action flag is not passed")
		t.Fatalf("Expected output: %s Got output: %s", wants, got)
	}
}

func TestInitRuns(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=init", "-p=./deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "Terraform has been successfully initialized!"
	if !strings.Contains(got, wants) {
		t.Fatalf("tf.Init() should run when action is 'init'")
	}
}

func TestPlanRuns(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=plan", "-p=deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "No changes."
	if !strings.Contains(got, wants) {
		t.Fatalf("tf.Plan() should run when action is 'plan'")
	}
}

func TestApplyRuns(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=apply", "-p=deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "Apply complete! Resources: 0 added, 0 changed, 0 destroyed"
	if !strings.Contains(got, wants) {
		t.Fatalf("tf.Plan() should run when action is 'plan'")
	}
}

func TestInvalidFlag(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=bad", "-p=./deploy"})
	if os.Getenv("TEST_DEPLOY_FAIL") == "true" {
		deployCmd.Run(deployCmd, []string{"deploy"})
		return
	}
	subCmd := exec.Command(os.Args[0], "-test.run=TestInvalidFlag")
	subCmd.Env = append(os.Environ(), "TEST_DEPLOY_FAIL=true")
	err := subCmd.Run()
	expectedErrorString := "exit status 1"
	if exiterr, ok := err.(*exec.ExitError); ok == true {
		if strings.Compare(expectedErrorString, exiterr.Error()) != 0 {
			t.Fatalf("deploy command should fail for invalid actions")
		}
	}
}

func TestInvalidPath(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=plan", "-p=./invalid"})
	if os.Getenv("TEST_DEPLOY_FAIL") == "true" {
		deployCmd.Run(deployCmd, []string{"deploy"})
		return
	}
	subCmd := exec.Command(os.Args[0], "-test.run=TestInvalidPath")
	subCmd.Env = append(os.Environ(), "TEST_DEPLOY_FAIL=true")
	err := subCmd.Run()
	expectedErrorString := "exit status 1"
	if exiterr, ok := err.(*exec.ExitError); ok == true {
		if strings.Compare(expectedErrorString, exiterr.Error()) != 0 {
			t.Fatalf("deploy command should fail for invalid path")
		}
	}
}
