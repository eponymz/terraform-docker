package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"tfd/cmd/deploy"
	"tfd/util"
)

func TestDefaultAction(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-p=./deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "Refreshing Terraform state in-memory prior to plan..."
	os.Chdir("../")
	if !strings.Contains(got, wants) {
		t.Fatalf("tf.Plan() should run as default when action flag is not passed")
	}
}

func TestInitRuns(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=init", "-p=./deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "Terraform has been successfully initialized!"
	os.Chdir("../")
	if !strings.Contains(got, wants) {
		t.Fatalf("tf.Init() should run when action is 'init'")
	}
}

func TestPlanRuns(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=plan", "-p=./deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "Refreshing Terraform state in-memory prior to plan..."
	os.Chdir("../")
	if !strings.Contains(got, wants) {
		t.Fatalf("tf.Plan() should run when action is 'plan'")
	}
}

func TestApplyRuns(t *testing.T) {
	deployCmd := deploy.GetCmd()
	deployCmd.ParseFlags([]string{"-w=default", "-a=apply", "-p=./deploy"})
	stdout, r, w := util.CaptureStdout()
	deployCmd.Run(deployCmd, []string{"deploy"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "Apply complete! Resources: 0 added, 0 changed, 0 destroyed"
	os.Chdir("../")
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
	subCmd := exec.Command(os.Args[0], "-test.run=TestPlanWithNonMappedFlag")
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
	subCmd := exec.Command(os.Args[0], "-test.run=TestPlanWithNonMappedFlag")
	subCmd.Env = append(os.Environ(), "TEST_DEPLOY_FAIL=true")
	err := subCmd.Run()
	expectedErrorString := "exit status 1"
	if exiterr, ok := err.(*exec.ExitError); ok == true {
		if strings.Compare(expectedErrorString, exiterr.Error()) != 0 {
			t.Fatalf("deploy command should fail for invalid actions")
		}
	}
}
