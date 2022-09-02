package bl

import (
	"errors"
	"fmt"
	"net/http"
	"parser_test/internal/models"
	"parser_test/internal/store/mongo"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type IParseHtml interface {
	ParseHtml(res *http.Response, selection models.Selection, saveFile bool) (map[string]interface{}, error)
}

func NewParseHtml(mongoRepo *mongo.MongoRepo) IParseHtml {
	return &parseHtml{mongoRepo: mongoRepo}
}

type parseHtml struct {
	mongoRepo *mongo.MongoRepo
}

func (ph *parseHtml) ParseHtml(res *http.Response, selection models.Selection, saveFile bool) (map[string]interface{}, error) {
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var parsArray = parseSelection(doc.Selection, selection.Find, "")

	result := make(map[string]interface{})
	// Find the review items
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		result[s.Find(".tg-1jpd").Text()] = s.Find(".tg-bbpb").Text()
	})

	if saveFile {
		saveResult := &models.ParserResult{
			Date:        time.Now(),
			Name:        res.Request.Host,
			Description: FromDto(parsArray),
		}
		err := ph.mongoRepo.Parser.Store(saveResult)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return result, nil
}

type ParseSelectionResult struct {
	Value      *string
	FoundValue *[]ParseSelectionResult
}

func FromDto(p *[]ParseSelectionResult) *[]models.ParseSelectionResult {
	var parsSelectResult []models.ParseSelectionResult

	if p != nil {
		for _, value := range *p {
			var convert = models.ParseSelectionResult{
				Value:      value.Value,
				FoundValue: FromDto(value.FoundValue),
			}
			parsSelectResult = append(parsSelectResult, convert)
		}
	}

	return &parsSelectResult
}

func parseSelection(doc *goquery.Selection, selection []*models.Find, find string) *[]ParseSelectionResult {
	var parsSelectResult []ParseSelectionResult

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
			var lineParse = ParseSelectionResult{}

			if startFindValue.GetValue {
				sText := s.Text()
				lineParse.Value = &sText
			}

			if startFindValue.Find != nil {
				lineParse.FoundValue = parseSelection(s, startFindValue.Find, find)
			}

			parsSelectResult = append(parsSelectResult, lineParse)
		})
	}

	return &parsSelectResult
}
