package direk

import (
	"errors"
	"regexp"

	"github.com/imroc/req/v3"
)

func Github(url string) (string, string, error) {
	releasesURL := regexp.MustCompile(`(?i)\bhttps?://.*github\.com/.*releases\S+`)
	match := releasesURL.FindString(url)
	if match == "" {
		return "", "", errors.New("no releases URL found")
	}

	client := req.C().SetRedirectPolicy(req.NoRedirectPolicy()) // Create a new client

	resp, err := client.R().Get(match)
	if err != nil {
		return "", "", errors.New("ERROR: " + err.Error())
	}
	defer resp.Body.Close()

	location, err := client.R().DisableAutoReadResponse().Get(match)
	if err != nil {
		return "", "", err
	}
	defer location.Body.Close()

	locationHeader := location.GetHeader("Location")
	if locationHeader != "" {
		return "", locationHeader, nil
	}
	return "", "", errors.New("Wrong Link it's suppose like this for example https://github.com/abcd/cdba/releases/download/v0.0.0.1/filename.apk")
}
