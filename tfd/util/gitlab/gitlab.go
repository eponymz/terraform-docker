package gitlab

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	gitlabBaseUrl    string = "https://gitlab.com/api/v4"
	pipelineApiToken string = os.Getenv("PIPELINE_API_TOKEN")
)

func PostMRComment(message string) {
	var (
		projectIID      string = os.Getenv("CI_MERGE_REQUEST_PROJECT_ID")
		mergeRequestIID string = os.Getenv("CI_MERGE_REQUEST_IID")
	)
	mrURL := fmt.Sprintf("%s/projects/%s/merge_requests/%s/notes", gitlabBaseUrl, projectIID, mergeRequestIID)
	fmt.Println("URL:>", mrURL)

	var jsonStr = []byte(fmt.Sprintf(`{"body":%q}`, message))
	req, err := http.NewRequest("POST", mrURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Private-Token", pipelineApiToken)
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
