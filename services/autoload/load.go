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

	for _, prefix := range []string{
		"https://www.mediafire.com",
		"https://krakenfiles.com/",
		"https://pixeldrain.com/",
		"https://gofile.io/",
		"https://wetransfer.com/",
		"https://we.tl/",
	} {
		if strings.HasPrefix(url, prefix) {
			switch prefix {
			case "https://www.mediafire.com":
				return direk.Mediafire(url)
			case "https://krakenfiles.com/":
				return direk.KrakenFiles(url)
			case "https://pixeldrain.com/":
				return direk.Pixeldrain(url)
			case "https://gofile.io/":
				return direk.Gofileio(url)
			case "https://wetransfer.com/":
				return direk.Wetransfer(url)
			case "https://we.tl/":
				return direk.Wetransfer(url)
			}
		}
	}

	return "", "", nil
}
