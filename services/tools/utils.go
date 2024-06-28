package tools

import "github.com/imroc/req/v3"

// GetRedirectUrl returns the redirect URL of the given URL.
//
// It takes a string parameter `url` which represents the URL to get the redirect URL from.
// It returns a string which represents the redirect URL.
func GetRedirectUrl(url string) string {
	client := req.C().SetRedirectPolicy(req.NoRedirectPolicy()) // Create a new client

	resp, err := client.R().Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	return resp.GetHeader("Location")
}
