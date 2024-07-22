package autoload

import (
	"errors"
	"regexp"
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
		"https://u.pcloud.link/",
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
			case "https://u.pcloud.link/":
				return direk.Pcloud(url)
			}
		}
	}
	if strings.HasPrefix(url, "https://github.com") {
		releasesURL := regexp.MustCompile(`(?i)\bhttps?://.*github\.com/[^/]+/[^/]+/releases/download/[^/]+/[^/]+$`)
		if !releasesURL.MatchString(url) {
			return "", "", errors.New("wrong Link it's suppose like this for example https://github.com/ViRb3/wgcf/releases/download/v2.2.22/wgcf_2.2.22_freebsd_386")
		}
		return direk.Github(url)
	}

	return "", "", nil
}
