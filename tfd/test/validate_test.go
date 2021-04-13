package test

import (
	"io/ioutil"
	"strings"
	"testing"
	"tfd/cmd/validate"
	"tfd/util"
)

func TestTfdocClean(t *testing.T) {
	tfdoc := validate.GettfdocCmd()
	stdout, r, w := util.CaptureStdout()
	tfdoc.Run(tfdoc, []string{"doc_clean"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "\n"
	if got != wants {
		t.Fatalf("Tfdoc wants %s, got %s", wants, got)
	}
}

func TestTfdocMismatched(t *testing.T) {
	tfdoc := validate.GettfdocCmd()
	stdout, r, w := util.CaptureStdout()
	tfdoc.Run(tfdoc, []string{"doc_mismatched"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "returned differences"
	if !strings.Contains(got, wants) {
		t.Fatalf("Tfdoc wants %s, got %s", wants, got)
	}
}

func TestTffmt(t *testing.T) {
	file := "fmt/fmt_test.tf"
	originalLintFile, _ := ioutil.ReadFile(file)

	tffmt := validate.GettffmtCmd()
	stdout, r, w := util.CaptureStdout()
	tffmt.Run(tffmt, []string{"fmt"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "fmt/fmt_test.t"
	if !strings.Contains(got, wants) {
		t.Fatalf("Tffmt wants %s, got %s", wants, got)
	}

	ioutil.WriteFile(file, originalLintFile, 744)
}

func TestTflint(t *testing.T) {
	tflint := validate.GettflintCmd()
	stdout, r, w := util.CaptureStdout()
	tflint.Run(tflint, []string{"lint"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "aws_instance_invalid_type"
	if !strings.Contains(got, wants) {
		t.Fatalf("Tflint wants %s, got %s", wants, got)
	}
}

func TestTfsec(t *testing.T) {
	tfsec := validate.GettfsecCmd()
	stdout, r, w := util.CaptureStdout()
	tfsec.Run(tfsec, []string{"fmt"})
	got := util.ReleaseStdout(stdout, r, w)
	wants := "potential problems detected"
	if !strings.Contains(got, wants) {
		t.Fatalf("Tfsec wants %s, got %s", wants, got)
	}
}
