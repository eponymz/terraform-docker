package test

import (
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
	tf "gitlab.com/edquity/devops/terraform-docker.git/tfd/util/terraform"
)

func TestInit(t *testing.T) {
	got := tf.Init("../tf_fail", )
	if got != 0 {
		t.Fatalf("tf.Init() should run and return 0, got: %d", got)
	}
}

func TestPlanFail(t *testing.T) {
	got := tf.Plan(".", "default")
	if got != 1 {
		t.Fatalf("tf.Plan() should return 1 when no configuration files are present, got: %d", got)
	}
}

func TestPlanEvalFail(t *testing.T) {
	got, _ := tf.PlanEval("../tf_fail/plan.tmp")
	wants := false
	if got != wants {
		t.Fatalf("tf.PlanEval() returned: %t - Expected: %t", got, wants)
	}
}

func TestApplyFail(t *testing.T) {
	got := tf.Apply(".", "default")
	if got != 1 {
		t.Fatalf("tf.Apply() should return 1 when no configuration files are present, got: %d", got)
	}
}

func TestWorkspaceValidDefault(t *testing.T) {
	got := tf.IsWorkspaceValid("default", ".")
	wants := false
	if got != wants {
		t.Fatalf("IsWorkspaceValid returned: %t. Expected: %t", got, wants)
	}
}

func TestWorkspaceValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		util.ExecExitCode("terraform", "workspace", "new", "valid")
		got := tf.IsWorkspaceValid("valid", ".")
		wants := true
		if got != wants {
			t.Fatalf("IsWorkspaceValid returned: %t. Expected: %t", got, wants)
		}
	})
	t.Run("Switch", func(t *testing.T) {
		got := tf.WorkspaceSwitch("valid", ".")
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
		got := tf.IsWorkspaceValid("valid", ".")
		wants := false
		if got != wants {
			t.Fatalf("IsWorkspaceValid returned: %t. Expected: %t", got, wants)
		}
	})
	t.Run("Switch", func(t *testing.T) {
		got := tf.WorkspaceSwitch("valid", ".")
		wants := false
		if got != wants {
			t.Fatalf("WorkspaceSwitch returned: %t. Expected: %t", got, wants)
		}
	})
}

func TestWorkspaceExecDefault(t *testing.T) {
	gotValid, gotVarFile := tf.WorkspaceExec("default", ".")
	wantsValid := true
	wantsVarFile := false
	if gotValid != wantsValid && gotVarFile != wantsVarFile {
		t.Fatalf("WorkspaceExec returned isValid: %t varFile: %t Expected isValid: %t varFile: %t", gotValid, gotVarFile, wantsValid, wantsVarFile)
	}
}
func TestWorkspaceExecValid(t *testing.T) {
	util.ExecExitCode("terraform", "workspace", "new", "valid")
	gotValid, gotVarFile := tf.WorkspaceExec("valid", ".")
	wantsValid := true
	wantsVarFile := true
	util.ExecExitCode("terraform", "workspace", "select", "default")
	util.ExecExitCode("terraform", "workspace", "delete", "valid")
	if gotValid != wantsValid && gotVarFile != wantsVarFile {
		t.Fatalf("WorkspaceExec returned isValid: %t varFile: %t Expected isValid: %t varFile: %t", gotValid, gotVarFile, wantsValid, wantsVarFile)
	}
}

func TestWorkspaceExecInvalid(t *testing.T) {
	gotValid, _ := tf.WorkspaceExec("valid", ".")
	wantsValid := false
	if gotValid != wantsValid {
		t.Fatalf("WorkspaceExec returned: %t Expected: %t with empty return string", gotValid, wantsValid)
	}
}
