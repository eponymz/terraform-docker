package test

import (
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
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
