package test

import (
	"testing"

	"gitlab.com/edquity/devops/terraform-docker.git/tfd/util"
)

func TestSliceContains(t *testing.T) {
	slice := []string{"some", "test", "values"}
	term := "values"
	if !util.SliceContains(slice, term) {
		t.Fatal("SliceContains(slice, term) wants true, got false.")
	}
}

func TestSliceEmpty(t *testing.T) {
	slice := []string{}
	if !util.SliceEmpty(slice) {
		t.Fatal("SliceEmpty(slice) wants true, got false.")
	}
}

func TestInExceptions(t *testing.T) {
	exceptions := []string{"this", "that"}
	term := "thats/so/raven"
	if !util.InExceptions(exceptions, term) {
		t.Fatal("InExceptions wants true, got false")
	}
}

func TestCleanString(t *testing.T) {
	testString := `  default
	* lab`
	wants := "default lab"
	got := util.CleanString(testString)
	if got != wants {
		t.Fatalf("CleanString got: %s, wants: %s", got, wants)
	}
}

func TestCleanStringLeadingStar(t *testing.T) {
	testString := `* default
	  lab`
	wants := "default lab"
	got := util.CleanString(testString)
	if got != wants {
		t.Fatalf("CleanString got: %s, wants: %s", got, wants)
	}
}

func TestCleanStringMultiLine(t *testing.T) {
	testString := `
	  default
	* lab
	`
	wants := "default lab"
	got := util.CleanString(testString)
	if got != wants {
		t.Fatalf("CleanString got: %s, wants: %s", got, wants)
	}
}

func TestCleanStringMulti(t *testing.T) {
	testString := `
	  default
	* lab
	  billing
	`
	wants := "default lab billing"
	got := util.CleanString(testString)
	if got != wants {
		t.Fatalf("CleanString got: %s, wants: %s", got, wants)
	}
}

func TestCleanStringSingle(t *testing.T) {
	testString := `* default
	`
	wants := "default"
	got := util.CleanString(testString)
	if got != wants {
		t.Fatalf("CleanString got: %s, wants: %s", got, wants)
	}
}
