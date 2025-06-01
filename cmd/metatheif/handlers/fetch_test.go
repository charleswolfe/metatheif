package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"metatheif/cmd/metatheif/tasks/fetcher"
)

var originalFetchURLs = fetcher.FetchURLs

func init() {
	fetcher.FetchURLs = func(urls []string) []fetcher.FetchResult {
		return []fetcher.FetchResult{
			{
				URL:    urls[0],
				Status: 200,
				Title:  "Test Title",
			},
		}
	}
}

func TestFetchHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/fetch", nil)
	w := httptest.NewRecorder()

	FetchHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", resp.StatusCode)
	}
}

func TestFetchHandler_BadPayload(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/fetch", bytes.NewBufferString("junk"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	FetchHandler(w, req)

	resp := w.Result()
	
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestFetchHandler_Success(t *testing.T) {
	body, _ := json.Marshal(map[string][]string{
		"urls": {"https://example.com"},
	})
	req := httptest.NewRequest(http.MethodPost, "/fetch", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	FetchHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var results []fetcher.FetchResult

	err := json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(results) != 1 || results[0].URL != "https://example.com" {
		t.Errorf("unexpected response: %+v", results)
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	// Restore original function
	fetcher.FetchURLs = originalFetchURLs
	
	os.Exit(code)
}
