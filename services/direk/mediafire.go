package direk

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
)

func Mediafire(url string) (string, string, error) {
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

	var href string
	doc2.Find("a.input.popsok[aria-label='Download file']").Each(func(i int, s *goquery.Selection) {
		href_exist, exist := s.Attr("href")
		if exist {
			href = href_exist
		} else {
			href = ""
		}
	})

	// finding filename
	var filename string
	doc2.Find("div.dl-btn-labelWrap > div.promoDownloadName > div.dl-btn-label").Each(func(i int, s *goquery.Selection) {
		text_filename, exist := s.Attr("title")
		if exist {
			filename = text_filename
		} else {
			filename = ""
		}
	})
	return filename, href, nil
}
