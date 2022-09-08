package parser_html_page

import (
	"errors"
	"fmt"
	"github.com/NordeN37/parser_html_page/models"
)

func GetResultParseHtml(parse models.Parse) (*[]models.ParseSelectionResult, error) {
	response, err := GetResponse(parse.Url, parse.HeaderSets)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", response.StatusCode, response.Status))
	}

	resultParse, err := ParseHtml(response.Body, parse.Selection)
	if err != nil {
		return nil, err
	}

	return resultParse, nil
}
