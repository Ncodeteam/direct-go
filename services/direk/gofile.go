package direk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path"
	"regexp"

	"github.com/imroc/req/v3"
)

func getToken(client *req.Client) (string, error) {
	resp, err := client.R().Post("https://api.gofile.io/accounts")
	if err != nil {
		return "", fmt.Errorf("guest account error: %v", err)
	}

	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("guest account changed again: %v", err)
	}

	// just get the "token" value
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return "", errors.New("data not found in guest account, maybe something goes wrong")
	}

	tokenValue, tokenExists := data["token"].(string)
	if !tokenExists {
		return "", errors.New("token not found, maybe something goes wrong")
	}
	return tokenValue, nil
}

func extractWebToken(input string, key string) (string, error) {
	re := regexp.MustCompile(key + `: "([^"]+)"`)
	match := re.FindStringSubmatch(input)
	if len(match) < 2 {
		return "", fmt.Errorf("key %s not found", key)
	}

	return match[1], nil
}

func getWebsiteToken(client *req.Client) (string, error) {
	// for getting websiteToken
	resp2, err := client.R().Get("https://gofile.io/dist/js/alljs.js")
	if err != nil {
		return "", fmt.Errorf("websiteToken Getter Broken: %v", err)
	}

	defer resp2.Body.Close()

	wtGet, err := io.ReadAll(resp2.Body)
	if err != nil {
		return "", err
	}

	wtValue, err := extractWebToken(string(wtGet), "wt")
	if err != nil {
		return "", err
	}
	return wtValue, nil
}

func Gofileio(link string) (string, string, error) {
	client := req.C() // Create a new client

	resp, err := client.R().Get(link)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// get the token first
	token, err := getToken(client)
	if err != nil {
		return "", "", err
	}

	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", "", fmt.Errorf("your file is invalid: %v", err)
	}

	idUrl := path.Base(parsedURL.Path)

	// get the file
	wtValue, _ := getWebsiteToken(client)

	url_dl := fmt.Sprintf("https://api.gofile.io/contents/%s?wt=%s", idUrl, wtValue)
	token_dl := fmt.Sprintf("Bearer %s", token)

	resp3, err := client.R().
		SetHeader("Authorization", token_dl).
		SetHeader("origin", "https://gofile.io").
		SetHeader("referer", "https://gofile.io/").
		Get(url_dl)
	if err != nil {
		return "", "", fmt.Errorf("gofile Changed again: %v", err)
	}

	defer resp3.Body.Close()

	// very hating moment about why i need to do this
	var response_dl map[string]interface{}
	if err := json.NewDecoder(resp3.Body).Decode(&response_dl); err != nil {
		return "", "", err
	}

	data, ok := response_dl["data"].(map[string]interface{})
	if !ok {
		return "", "", errors.New("data not found in json")
	}

	children, ok := data["children"].(map[string]interface{})
	if !ok {
		return "", "", errors.New("children not found in json")
	}

	var file_name, link_dl string

	for _, child := range children {
		childMap := child.(map[string]interface{})
		file_name = childMap["name"].(string)
		link_dl = childMap["link"].(string)
		break
	}
	return file_name, link_dl, nil
}
