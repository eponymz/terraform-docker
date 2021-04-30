package test

import (
	"strings"
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
)

func TestExecExceptAll(t *testing.T) {
	exceptions := []string{"test"}
	got := util.ExecExcept(exceptions, "ls", "test")
	wants := ""
	if got != wants {
		t.Fatalf("ExecExcept wants %s got %s", wants, got)
	}
}

func TestExecExceptOther(t *testing.T) {
	exceptions := []string{"bad"}
	got := util.ExecExcept(exceptions, "ls", ".")
	wants := "exec_test.go"
	if !strings.Contains(got, wants) {
		t.Fatalf("ExecExcept wants %s got %s", wants, got)
	}
}

func TestExecExceptNone(t *testing.T) {
	exceptions := []string{""}
	got := util.ExecExcept(exceptions, "ls", ".")
	wants := "exec_test.go"
	if !strings.Contains(got, wants) {
		t.Fatalf("ExecExcept wants %s got %s", wants, got)
	}
}

func TestExecExceptRPositive(t *testing.T) {
	exceptions := []string{"../test"}
	got := util.ExecExceptR(exceptions, "ls", "..")
	wants := "main.go"
	if !strings.Contains(got, wants) {
		t.Fatalf("ExecExceptR wants %s got %s", wants, got)
	}
}

func TestExecExceptRNegative(t *testing.T) {
	exceptions := []string{"../test", "empty"}
	got := util.ExecExceptR(exceptions, "ls", "..")
	notWanted := "exec_test.go"
	if strings.Contains(got, notWanted) {
		t.Fatalf("ExecExceptR return should not contain %s, got %s", notWanted, got)
	}
}

func TestExecExceptRfmt(t *testing.T) {
	exceptions := []string{"../test", "empty"}
	got := util.ExecExceptR(exceptions, "terraform fmt", ".", "--help")
	wants := "Usage: terraform fmt"
	if !strings.Contains(got, wants) {
		t.Fatalf("ExecExceptR wants %s got %s", wants, got)
	}
}

func TestExecExitCodeSuccess(t *testing.T) {
	got := util.ExecExitCode("ls", ".")
	wants := 0
	if got != wants {
		t.Fatalf("ExecExitCode wants %d got %d", wants, got)
	}
}

func TestExecExitCodeNoCommand(t *testing.T) {
	got := util.ExecExitCode("lolcat", "--help")
	wants := 1
	if got != wants {
		t.Fatalf("ExecExitCode should fail to run commands that do not exist. Got %d, wants %d", got, wants)
	}
}
