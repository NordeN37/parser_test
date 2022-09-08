package bl

import (
	pmodels "github.com/NordeN37/parser_html_page/models"
	"parser_test/internal/models"
	"parser_test/internal/store/mongo"
	"time"
)

type IParseHtml interface {
	ParseHtml(parseResult *[]pmodels.ParseSelectionResult, hostname string, saveFile bool) error
}

func NewParseHtml(mongoRepo *mongo.MongoRepo) IParseHtml {
	return &parseHtml{mongoRepo: mongoRepo}
}

type parseHtml struct {
	mongoRepo *mongo.MongoRepo
}

func (ph *parseHtml) ParseHtml(parseResult *[]pmodels.ParseSelectionResult, hostname string, saveFile bool) error {
	var parsArray = fromDto(parseResult)

	if saveFile {
		saveResult := &models.ParserResult{
			Date:        time.Now(),
			Name:        hostname,
			Description: parsArray,
		}
		err := ph.mongoRepo.Parser.Store(saveResult)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func fromDto(p *[]pmodels.ParseSelectionResult) *[]models.ParseSelectionResult {
	var parsSelectResult []models.ParseSelectionResult

	if p != nil {
		for _, value := range *p {
			var convert = models.ParseSelectionResult{
				Value:      value.Value,
				FoundValue: fromDto(value.FoundValue),
			}
			parsSelectResult = append(parsSelectResult, convert)
		}
	}

	return &parsSelectResult
}
