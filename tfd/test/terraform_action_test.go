package test

import (
	"testing"
	"tfd/util"
	tf "tfd/util/terraform"
)

func TestInit(t *testing.T) {
	got := tf.Init("../tf_fail")
	if got != 0 {
		t.Fatalf("tf.Init() should run and return 0, got: %d", got)
	}
}

func TestPlanFail(t *testing.T) {
	got := tf.Plan("../tf_fail", "default")
	if got != 1 {
		t.Fatalf("tf.Plan() should return 1 when no configuration files are present, got: %d", got)
	}
}

func TestApplyFail(t *testing.T) {
	got := tf.Apply("../tf_fail", "default")
	if got != 1 {
		t.Fatalf("tf.Apply() should return 1 when no configuration files are present, got: %d", got)
	}
}

func TestWorkspaceValidDefault(t *testing.T) {
	got := tf.IsWorkspaceValid("default")
	wants := false
	if got != wants {
		t.Fatalf("IsWorkspaceValid returned: %t. Expected: %t", got, wants)
	}
}

func TestWorkspaceValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		util.ExecExitCode("terraform", "workspace", "new", "valid")
		got := tf.IsWorkspaceValid("valid")
		wants := true
		if got != wants {
			t.Fatalf("IsWorkspaceValid returned: %t. Expected: %t", got, wants)
		}
	})
	t.Run("Switch", func(t *testing.T) {
		got := tf.WorkspaceSwitch("valid")
		wants := true
		util.ExecExitCode("terraform", "workspace", "select", "default")
		util.ExecExitCode("terraform", "workspace", "delete", "valid")
		if got != wants {
			t.Fatalf("WorkspaceSwitch returned: %t. Expected: %t", got, wants)
		}
	})
}

func TestWorkspaceInvalid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		got := tf.IsWorkspaceValid("valid")
		wants := false
		if got != wants {
			t.Fatalf("IsWorkspaceValid returned: %t. Expected: %t", got, wants)
		}
	})
	t.Run("Switch", func(t *testing.T) {
		got := tf.WorkspaceSwitch("valid")
		wants := false
		if got != wants {
			t.Fatalf("WorkspaceSwitch returned: %t. Expected: %t", got, wants)
		}
	})
}

func TestWorkspaceExecDefault(t *testing.T) {
	gotValid, gotString := tf.WorkspaceExec("default")
	wantsValid := true
	if gotValid != wantsValid && gotString != "" {
		t.Fatalf("WorkspaceExec returned: %t with return string %s Expected: %t with empty return string", gotValid, gotString, wantsValid)
	}
}
func TestWorkspaceExecValid(t *testing.T) {
	util.ExecExitCode("terraform", "workspace", "new", "valid")
	gotValid, gotString := tf.WorkspaceExec("valid")
	wantsValid := true
	util.ExecExitCode("terraform", "workspace", "select", "default")
	util.ExecExitCode("terraform", "workspace", "delete", "valid")
	if gotValid != wantsValid && gotString != " -var-file=valid.tfvars" {
		t.Fatalf("WorkspaceExec returned: %t with return string %s Expected: %t with empty return string", gotValid, gotString, wantsValid)
	}
}

func TestWorkspaceExecInvalid(t *testing.T) {
	gotValid, _ := tf.WorkspaceExec("valid")
	wantsValid := false
	if gotValid != wantsValid {
		t.Fatalf("WorkspaceExec returned: %t Expected: %t with empty return string", gotValid, wantsValid)
	}
}
