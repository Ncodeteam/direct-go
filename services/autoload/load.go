package autoload

import (
	"errors"
	"strings"

	"github.com/Ncodeteam/direct-go/services/direk"
)

func DirectLink(url string) (string, string, error) {
	if url == "" {
		return "", "", errors.New("url is empty")
	}

	if strings.HasPrefix(url, "https://www.mediafire.com/file/") {
		filename, href, err := direk.Mediafire(url)
		if err != nil {
			return "", "", err
		}
		return filename, href, err
	}

	return "", "", nil
}
