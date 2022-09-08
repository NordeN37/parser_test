package parser_controller

import (
	"encoding/json"
	"github.com/NordeN37/parser_html_page"
	pmodels "github.com/NordeN37/parser_html_page/models"
	"log"
	"net/http"
	"net/url"
)

func (pc *ParserController) Parser(_ http.ResponseWriter, r *http.Request) (interface{}, error) {
	var parse pmodels.Parse
	if err := json.NewDecoder(r.Body).Decode(&parse); err != nil {
		return nil, err
	}

	resultParse, err := parser_html_page.GetResultParseHtml(parse)
	if err != nil {
		return nil, err
	}

	var collectionHostName string
	url, err := url.Parse(parse.Url)
	if err != nil {
		collectionHostName = "host_not_found"
		log.Println(err.Error())
	}
	collectionHostName = url.Host

	err = pc.bl.Parser.ParseHtml(resultParse, collectionHostName, true)
	if err != nil {
		return nil, err
	}

	return resultParse, nil
}
