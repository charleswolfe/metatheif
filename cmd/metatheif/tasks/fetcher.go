package fetcher

import (
	"log"
	"fmt"
	"net/http"
	"bytes"
	"io"
	"github.com/PuerkitoBio/goquery"
)

// Package fetcher provides functionality to fetch URLs and return their status and title.
type FetchRequest struct {
	URL string `json:"url"`
}

// FetchResult represents the result of fetching a URL, including its status and title.
type FetchResult struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// FetchURLs takes a slice of URLs and returns a slice of FetchResult.
func FetchURLs(urls []string) []FetchResult {
	results := make([]FetchResult, 0)
	ch := make(chan FetchResult)

	for _, url := range urls {
		go func(u string) {
			resp, err := http.Get(u)

			if err != nil {
				log.Printf("failed to fetch %s: %v\n", u, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)

			//parse body
			title, err := parseHtmlBody(body, "title")

			ch <- FetchResult{
				URL:    u,
				Status: resp.StatusCode,
				Title:  title,
			}



		}(url)
	}

	for range urls {
		result := <-ch
		results = append(results, result)
	}

	return results
}

func parseHtmlBody(body []byte, searchTerm string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	search := doc.Find(searchTerm).Text()
	return search, nil
}
