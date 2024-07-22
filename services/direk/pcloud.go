package direk

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/imroc/req/v3"
)

func Pcloud(url string) (string, string, error) {
	client := req.C() // Create a new client

	resp, err := client.R().Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// parse the response body
	body := resp.String()

	// Define regex patterns
	metadataNamePattern := `"name"\s*:\s*"([^"]+)"`
	downloadLinkPattern := `"downloadlink"\s*:\s*"([^"]+)"`

	// Compile regex patterns
	metadataNameRegex := regexp.MustCompile(metadataNamePattern)
	downloadLinkRegex := regexp.MustCompile(downloadLinkPattern)

	// Find matches in the body
	filenameResult := metadataNameRegex.FindStringSubmatch(body)
	downloadLinkResult := downloadLinkRegex.FindStringSubmatch(body)

	// Check if matches were found
	if len(filenameResult) < 2 || len(downloadLinkResult) < 2 {
		fmt.Println("Unable to extract filename or download link")
		return "", "", nil
	}

	// Extract filename and download link
	filename := filenameResult[1]
	downloadLink := strings.Replace(downloadLinkResult[1], "\\/", "/", -1)

	return filename, downloadLink, nil
}
