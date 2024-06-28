package direk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
)

func KrakenFiles(url string) (string, string, error) {
	client := req.C() // Create a new client

	resp, err := client.R().Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// for finding
	doc2, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}

	// finding action and token
	var token string
	var post_url string
	var filename string

	doc2.Find("form#dl-form").Each(func(i int, s *goquery.Selection) {
		action_url, _ := s.Attr("action")
		if action_url == "" {
			err = errors.New("action is empty")
		} else {
			post_url = "https://krakenfiles.com" + action_url
		}
	})
	var tokenCheck string
	doc2.Find("input[name='token']").Each(func(i int, s *goquery.Selection) {
		tokenCheck, _ = s.Attr("value")
		if tokenCheck == "" {
			err = errors.New("token is empty")
		} else {
			token = tokenCheck
		}
	})

	resp, err = client.R().SetFormData(map[string]string{"token": token}).Post(post_url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// for href
	var href string
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", fmt.Errorf("error decoding response: %v", err)
	}
	status, ok := response["status"].(string)
	if !ok || status != "ok" {
		return "", "", nil
	}
	href, ok = response["url"].(string)
	if !ok {
		return "", "", errors.New("url is not string")
	}

	// for filename
	doc2.Find("div.coin-info > span.coin-name > h5").Each(func(i int, s *goquery.Selection) {
		text_filename := s.Text()
		if text_filename != "" {
			filename = text_filename
		} else {
			filename = ""
		}
	})
	return filename, href, err

}
