package gitlab

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	gitlabBaseUrl string = "https://gitlab.com/api/v4"
)

func PostMRComment(message string) error {
	logrus.Info("PostMRComment running")
	var (
		returnErr error = nil
		jsonStr         = []byte(fmt.Sprintf(`{"body":%q}`, message))
	)
	pipelineApiToken, tokenOk := os.LookupEnv("PIPELINE_API_TOKEN")
	projectIID, projectIdOk := os.LookupEnv("CI_MERGE_REQUEST_PROJECT_ID")
	mergeRequestIID, mrIdOk := os.LookupEnv("CI_MERGE_REQUEST_IID")

	if !tokenOk || !projectIdOk || !mrIdOk {
		errMsg := fmt.Sprintf("some or all required variables not present - tokenOk: %t projectIdOk: %t mrIdOk: %t", tokenOk, projectIdOk, mrIdOk)
		returnErr = errors.New(errMsg)
	} else {
		mrURL := fmt.Sprintf("%s/projects/%s/merge_requests/%s/notes", gitlabBaseUrl, projectIID, mergeRequestIID)
		logrus.Tracef("Request URL: %s", mrURL)

		req, err := http.NewRequest("POST", mrURL, bytes.NewBuffer(jsonStr))
		req.Header.Set("Private-Token", pipelineApiToken)
		req.Header.Set("cache-control", "no-cache")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logrus.Errorf("PostMRComment HTTP Call failed: %s", err)
			returnErr = err
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			logrus.Warnf("Response Status: %s", resp.Status)
		} else {
			logrus.Infof("Response Status: %s", resp.Status)
		}
	}
	return returnErr
}
