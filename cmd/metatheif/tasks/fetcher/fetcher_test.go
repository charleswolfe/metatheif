package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchURLs(t *testing.T) {
	fakePageText := `<html><head><title>Test Page</title></head><body>Hello</body></html>`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakePageText))
	}))

	defer mockServer.Close()

	// Prepare input and expected output
	urls := []string{mockServer.URL}
	results := FetchURLs(urls)

	// Basic checks
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	result := results[0]
	if result.Status != 200 {
		t.Errorf("expected status 200, got %d", result.Status)
	}
	if result.Title != "Test Page" {
		t.Errorf("expected title 'Test Page', got '%s'", result.Title)
	}
	if result.URL != mockServer.URL {
		t.Errorf("expected URL %s, got %s", mockServer.URL, result.URL)
	}
}
