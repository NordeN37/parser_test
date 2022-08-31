package parser_controller

import (
	"encoding/json"
	"net/http"
	"parser_test/internal/models"
	"parser_test/internal/util"
)

func (pc *ParserController) Parser(_ http.ResponseWriter, r *http.Request) (interface{}, error) {
	var parse models.Parse
	if err := json.NewDecoder(r.Body).Decode(&parse); err != nil {
		return nil, err
	}

	response, err := util.GetResponse(parse.Url, parse.HeaderSets)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resultParse, err := pc.bl.Parser.ParseHtml(response, parse.Selection, true)
	if err != nil {
		return nil, err
	}

	return resultParse, nil
}
