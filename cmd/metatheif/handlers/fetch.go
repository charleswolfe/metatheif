package handlers

// The handler will decode the request, validate it, and then call the fetcher to process the URLs.
import (
	"encoding/json"
	"net/http"
	"metatheif/cmd/metatheif/tasks"
)

// It expects a POST request with a JSON body containing an array of URLs.
type requestPayload struct {
	URLs []string `json:"urls"`
}

// FetchHandler handles the /fetch endpoint for fetching URLs
func FetchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Site is working, please use POST method only", http.StatusMethodNotAllowed)
		return
	}

	var payload requestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil || len(payload.URLs) == 0 {
		http.Error(w, "The payload is junk!", http.StatusBadRequest)
		return
	}

	results := fetcher.FetchURLs(payload.URLs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}