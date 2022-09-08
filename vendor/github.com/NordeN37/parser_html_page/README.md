# parser_html_page

## Installation

    $ go get github.com/NordeN37/parser_html_page

## Methods
returns the response from the given URL.
GetResponse(url string, headerSets []*models.HeaderSet) (*http.Response, error)

It uses goquery at its core. Parses the document and recursively checks the Selection
ParseHtml(~)

The general method accepts a link to the page, headers, selectors.
Makes a request, receives a document, and requests
GetResultParseHtml(~)

## Examples

```Go
package main

import (
	"encoding/json"
	"github.com/NordeN37/parser_html_page"
	pmodels "github.com/NordeN37/parser_html_page/models"
	"log"
)

func main() {
	var parse pmodels.Parse
	err := json.Unmarshal([]byte(data), &parse)
	if err != nil {
		log.Fatal(err)
	}

	result, err := parser_html_page.GetResultParseHtml(parse)
	if err != nil {
		log.Fatal(err)
	}

	byteResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(byteResult))
}

var data = `{
  "URL": "https://example.com/",
  "HeaderSets": [
    {
      "Key": "User-Agent",
      "Value": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0"
    },
    {
      "Key": "Accept",
      "Value": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
    },
    {
      "Key": "Content-Type",
      "Value": "text/html; charset=utf-8"
    }
  ],
  "Selection": {
    "Find": [
      {
        "Tag": "div",
        "GetValue": false,
        "Find": [
          {
            "Tag": "h1",
            "GetValue": true
          },
          {
            "Tag": "p",
            "GetValue": true
          }
        ]
      }
    ]
  }
}`

```