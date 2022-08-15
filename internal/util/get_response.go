package util

import (
	"net/http"
	"parser_test/internal/models"
)

// GetResponse returns the response from the given URL.
func GetResponse(url string, headerSets []*models.HeaderSet) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	switch {
	case headerSets != nil:
		for _, set := range headerSets {
			if set.Key != "" && set.Value != "" {
				req.Header.Set(set.Key, set.Value)
			}
		}
	default:
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Set("Content-Type", "text/html; charset=utf-8")
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
