package parser_controller

import (
	"encoding/json"
	"net/http"
	"parser_test/internal/bl"
	"parser_test/internal/models"
	"parser_test/internal/util"
)

func Parser(_ http.ResponseWriter, r *http.Request) (interface{}, error) {
	var parse models.Parse
	if err := json.NewDecoder(r.Body).Decode(&parse); err != nil {
		return nil, err
	}

	response, err := util.GetResponse(parse.Url, parse.HeaderSets)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resultParse, err := bl.ParseHtml(response, parse.Selection, true)
	if err != nil {
		return nil, err
	}

	return resultParse, nil
}
