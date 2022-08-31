package bl

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"parser_test/internal/models"
	"parser_test/internal/store/mongo"
	"time"
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

	//var parsSelectResult = new(parseSelectionResult)
	//parseSelection(doc, parsSelectResult, selection.Find, "")

	result := make(map[string]interface{})
	// Find the review items
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		result[s.Find(".tg-1jpd").Text()] = s.Find(".tg-bbpb").Text()
	})

	if saveFile {
		saveResult := &models.ParserResult{
			Date:        time.Now(),
			Name:        res.Request.Host,
			Description: result,
		}
		err := ph.mongoRepo.Parser.Store(saveResult)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return result, nil
}

type parseSelectionResult struct {
	value      *string
	foundValue *parseSelectionResult
}

func parseSelection(doc *goquery.Document, result *parseSelectionResult, selection []*models.Find, find string) {
	for _, startFindValue := range selection {
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
			switch {
			case startFindValue.Find != nil && startFindValue.GetValue:
				sText := s.Text()
				result.value = &sText
				parseSelection(doc, result.foundValue, startFindValue.Find, find)
			case startFindValue.GetValue:
				sText := s.Text()
				result.value = &sText
			case startFindValue.Find != nil:
				parseSelection(doc, result.foundValue, startFindValue.Find, find)
			}
		})
	}
	return
}
