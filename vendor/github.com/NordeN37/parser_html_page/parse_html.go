package parser_html_page

import (
	"github.com/NordeN37/parser_html_page/models"
	"github.com/PuerkitoBio/goquery"
	"io"
)

func ParseHtml(res io.ReadCloser, selection models.Selection) (*[]models.ParseSelectionResult, error) {
	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		return nil, err
	}

	return parseSelection(doc.Selection, selection.Find), nil
}

func parseSelection(doc *goquery.Selection, selection []*models.Find) *[]models.ParseSelectionResult {
	var parsSelectResult []models.ParseSelectionResult

	for _, startFindValue := range selection {
		var find string
		if startFindValue.Tag != nil {
			find = *startFindValue.Tag
		}
		if startFindValue.Class != nil {
			find += " " + *startFindValue.Class
		}
		if startFindValue.Id != nil {
			find += " " + *startFindValue.Id
		}

		// Find the review items
		doc.Find(find).Each(func(i int, s *goquery.Selection) {
			var lineParse = models.ParseSelectionResult{}

			if startFindValue.GetValue {
				sText := s.Text()
				lineParse.Value = &sText
			}

			if startFindValue.Find != nil {
				lineParse.FoundValue = parseSelection(s, startFindValue.Find)
			}

			parsSelectResult = append(parsSelectResult, lineParse)
		})
	}

	return &parsSelectResult
}
