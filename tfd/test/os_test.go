package test

import (
	"os"
	"testing"
	"tfd/util"
)

func TestDirTreeListOneUp(t *testing.T) {
	got := util.DirTreeList("..")
	wants := "../test"
	if !util.SliceContains(got, wants) {
		t.Fatalf("DirTreeList wants %s got %s", wants, got)
	}
}

func TestDirTreeListPwd(t *testing.T) {
	got := util.DirTreeList(".")
	notWanted := "exec_test.go"
	if util.SliceContains(got, notWanted) {
		t.Fatalf("DirTreeList return should not contain %s, got %s", notWanted, got)
	}
}

func TestSafeChangeDir(t *testing.T) {
	got := util.SafeChangeDir("./deploy")
	os.Chdir(("../"))
	if got != nil {
		t.Fatalf("util.SafeChangeDir should return 'nil'. Got: %v", got)
	}
}

func TestSafeChangeDirFail(t *testing.T) {
	got := util.SafeChangeDir("./invalid")
	if got == nil {
		t.Fatalf("util.SafeChangeDir should return an error. Got: %s", got)
	}
}
