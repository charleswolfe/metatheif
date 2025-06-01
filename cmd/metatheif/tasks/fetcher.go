package fetcher

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

			// return something for now
			ch <- FetchResult{
				URL:    u,
				Status: 200, // Placeholder
				Title:  "Not implemented",
			}
		}(url)
	}

	for range urls {
		result := <-ch
		results = append(results, result)
	}

	return results
}
