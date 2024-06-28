package direk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Ncodeteam/direct-go/services/tools"
	"github.com/imroc/req/v3"
)

func Wetransfer(url string) (string, string, error) {
	// check fileid
	client := req.C() // Create a new client

	var link string
	if strings.HasPrefix(url, "https://we.tl/") {
		link = tools.GetRedirectUrl(url)
	} else {
		link = url
	}
	parts := strings.Split(link, "/")
	fileID := parts[len(parts)-1]
	if parts == nil || fileID == "" {
		return "", "", errors.New("fileID not found")
	}

	jsonData := map[string]string{
		"security_hash": fileID,
		"intent":        "entire_transfer",
	}

	resp_post, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonData).
		Post(fmt.Sprintf("https://wetransfer.com/api/v4/transfers/%s/download", parts[len(parts)-2]))
	if err != nil {
		return "", "", err
	}
	defer resp_post.Body.Close()

	// request filename
	jsonDataFileName := map[string]string{
		"security_hash": fileID,
	}

	resp_filename, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonDataFileName).
		Post(fmt.Sprintf("https://wetransfer.com/api/v4/transfers/%s/prepare-download", parts[len(parts)-2]))
	if err != nil {
		return "", "", err
	}
	defer resp_filename.Body.Close()

	var response_filename map[string]interface{}
	if err := json.NewDecoder(resp_filename.Body).Decode(&response_filename); err != nil {
		return "", "", err
	}
	var filename string
	for _, item := range response_filename["items"].([]interface{}) {
		filename = item.(map[string]interface{})["name"].(string)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp_post.Body).Decode(&result)
	if err != nil {
		return "", "", err
	}

	if directLink, ok := result["direct_link"].(string); ok {
		return filename, directLink, nil
	} else if message, ok := result["message"].(string); ok {
		return "", "", fmt.Errorf("ERROR: %s", message)
	} else if errorMsg, ok := result["error"].(string); ok {
		return "", "", fmt.Errorf("ERROR: %s", errorMsg)
	} else {
		return "", "", fmt.Errorf("ERROR: cannot find direct link")
	}

}
