package util

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// GetGoQueryDoc fetches an URL and parses the document to GoQuery.
func GetGoQueryDoc(url string) (*goquery.Document, error) {
	var (
		err  error
		resp *http.Response
	)

	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Convert HTML into goquery document
	return goquery.NewDocumentFromReader(resp.Body)
}
