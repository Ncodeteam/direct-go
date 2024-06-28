package services

import "errors"

func direct_link(url string) (string, error) {
	if url == "" {
		return "", errors.New("url is empty")
	}

	return url, nil
}
