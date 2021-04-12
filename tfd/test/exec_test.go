package test

import (
	"os"
	"strings"
	"testing"
	"tfd/cmd"
	"tfd/util"
)

func TestMain(m *testing.M) {
	cmd.InitConfig() // Unified logging configuration from app
	code := m.Run()
	os.Exit(code)
}

func TestSliceContains(t *testing.T) {
	slice := []string{"some", "test", "values"}
	term := "values"
	if !util.SliceContains(slice, term) {
		t.Fatal("SliceContains(slice, term) wants true, got false.")
	}
}

func TestExecExcept(t *testing.T) {
	exceptions := []string{"test"}
	got := util.ExecExcept(exceptions, "ls", "test")
	wants := ""
	if got != wants {
		t.Fatalf("ExecExcept wants %s got %s", wants, got)
	}
}

func TestExecExcept2(t *testing.T) {
	exceptions := []string{"bad"}
	got := util.ExecExcept(exceptions, "ls", ".")
	wants := "exec_test.go"
	if !strings.Contains(got, wants) {
		t.Fatalf("ExecExcept wants %s got %s", wants, got)
	}
}

func TestDirTreeList(t *testing.T) {
	got := util.DirTreeList("..")
	wants := "../test"
	if !util.SliceContains(got, wants) {
		t.Fatalf("DirTreeList wants %s got %s", wants, got)
	}
}

func TestDirTreeList2(t *testing.T) {
	got := util.DirTreeList(".")
	notWanted := "exec_test.go"
	if util.SliceContains(got, notWanted) {
		t.Fatalf("DirTreeList return should not contain %s, got %s", notWanted, got)
	}
}

func TestExecExceptR(t *testing.T) {
	exceptions := []string{"../test", "empty"}
	got := util.ExecExceptR(exceptions, "ls", "..")
	wants := "main.go"
	if !strings.Contains(got, wants) {
		t.Fatalf("ExecExceptR wants %s got %s", wants, got)
	}
}

func TestExecExceptR2(t *testing.T) {
	exceptions := []string{"../test", "empty"}
	got := util.ExecExceptR(exceptions, "ls", "..")
	notWanted := "exec_test.go"
	if strings.Contains(got, notWanted) {
		t.Fatalf("ExecExceptR return should not contain %s, got %s", notWanted, got)
	}
}
