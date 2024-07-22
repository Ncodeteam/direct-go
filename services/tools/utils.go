package tools

import (
	"fmt"
	"io"
	"os"

	"github.com/imroc/req/v3"
)

// Testresponse writes the response body to a file named "result.txt".
// It returns an error if any.
func Testresponse(url string) error {
	client := req.C().SetRedirectPolicy(req.NoRedirectPolicy()) // Create a new client

	resp, err := client.R().Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // read the response body
	if err != nil {
		return err // return an empty string and the error
	}
	filename := "result.txt"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		return err
	}
	fmt.Println("Successfully saved to", filename)
	return nil
}

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

// GetStreamUrl returns the stream URL of the given URL.
//
// It takes a string parameter `url` which represents the URL to get the stream URL from.
// It returns a string which represents the stream URL.
//
// The function creates a new client using `req.C().SetRedirectPolicy(req.NoRedirectPolicy())`.
// It sends a GET request to the given `url` using `client.R().DisableAutoReadResponse().Get(url)`.
// If there is an error, it returns an empty string.
// It closes the response body using `defer resp.Body.Close()`.
// Finally, it returns the value of the "Location" header from the response.
func GetStreamUrl(url string) string {

	client := req.C().SetRedirectPolicy(req.NoRedirectPolicy()) // Create a new client

	resp, err := client.R().DisableAutoReadResponse().Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	return resp.GetHeader("Location")
}
