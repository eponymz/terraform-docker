package test

import (
	"fmt"
	"os"
	"testing"
	"tfd/util/gitlab"
)

func TestVarsNotSet(t *testing.T) {
	t.Run("Token Not Set", func (t *testing.T) {
		os.Unsetenv("PIPELINE_API_TOKEN")
		os.Setenv("CI_MERGE_REQUEST_PROJECT_ID", "1")
		os.Setenv("CI_MERGE_REQUEST_IID", "1")
		wants := "some or all required variables not present - tokenOk: false projectIdOk: true mrIdOk: true"
		got := gitlab.PostMRComment("TestTokenNotSet")
		if fmt.Sprintf("%s", got) != fmt.Sprintf("%s", wants) {
			t.Fatalf(`
			Got: %s
			Wants: %s`, got, wants)
		}
	})
	t.Run("Project ID Not Set", func (t *testing.T) {
		os.Setenv("PIPELINE_API_TOKEN", "")
		os.Unsetenv("CI_MERGE_REQUEST_PROJECT_ID")
		os.Setenv("CI_MERGE_REQUEST_IID", "1")
		wants := "some or all required variables not present - tokenOk: true projectIdOk: false mrIdOk: true"
		got := gitlab.PostMRComment("TestTokenNotSet")
		if fmt.Sprintf("%s", got) != fmt.Sprintf("%s", wants) {
			t.Fatalf(`
			Got: %s
			Wants: %s`, got, wants)
		}
	})
	t.Run("Merge Request ID Not Set", func (t *testing.T) {
		os.Setenv("PIPELINE_API_TOKEN", "")
		os.Setenv("CI_MERGE_REQUEST_PROJECT_ID", "1")
		os.Unsetenv("CI_MERGE_REQUEST_IID")
		wants := "some or all required variables not present - tokenOk: true projectIdOk: true mrIdOk: false"
		got := gitlab.PostMRComment("TestTokenNotSet")
		if fmt.Sprintf("%s", got) != fmt.Sprintf("%s", wants) {
			t.Fatalf(`
			Got: %s
			Wants: %s`, got, wants)
		}
	})
}

func TestFailedComment(t *testing.T) {
	os.Setenv("PIPELINE_API_TOKEN", "")
	os.Setenv("CI_MERGE_REQUEST_PROJECT_ID", "25611223")
	os.Setenv("CI_MERGE_REQUEST_IID", "11")
	if got := gitlab.PostMRComment("TestTokenNotSet"); got != nil {
		t.Fatalf("PostMRComment should return <nil> when all values are set. Got: %v", got)
	}
}
