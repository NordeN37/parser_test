package bl

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"parser_test/internal/models"
	"parser_test/internal/util"
)

func ParseHtml(res *http.Response, selection models.Selection, saveFile bool) (map[string]interface{}, error) {
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	// Find the review items
	doc.Find(".media-block__content").Each(func(i int, s *goquery.Selection) {
		result[s.Find("H4").Text()] = s.Find(".link-service").Text()
	})

	if saveFile {
		if err = util.CreateFileResultJson(result, "./result/", res.Request.Host); err != nil {
			return nil, err
		}
		return nil, nil
	}
	return result, nil
}
