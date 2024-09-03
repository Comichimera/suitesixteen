package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	githubAPIURL = "https://api.github.com"
)

type GitHubFile struct {
	Message string `json:"message"`
	Content string `json:"content"`
}

func UploadFileToGitHub(owner, repo, path, token, message, filePath string) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	content := base64.StdEncoding.EncodeToString(fileContent)

	file := GitHubFile{
		Message: message,
		Content: content,
	}

	fileJSON, err := json.Marshal(file)
	if err != nil {
		return fmt.Errorf("could not marshal json: %v", err)
	}

	url := fmt.Sprintf("%s/repos/%s/%s/contents/%s", githubAPIURL, owner, repo, path)
	fmt.Printf("Attempting to upload to URL: %s\n", url)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(fileJSON))
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 404 {
		return fmt.Errorf("404 Not Found. Please check the following:\n- Repository owner: '%s'\n- Repository name: '%s'\n- Path: '%s'\n- Branch: Is the branch correct?\n\nResponse: %s", owner, repo, path, string(body))
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("failed to upload file, status: %d, response: %s", resp.StatusCode, string(body))
	}

	fmt.Println("File uploaded successfully!")
	return nil
}
