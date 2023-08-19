package friendtech

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func GetRecentlyJoinedv2() (*Response, error) {
	url := "https://friendmex.com/api/stats/newest"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	var result Response

	// Decoding directly from the HTTP response body into the Response object
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %v", err)
	}

	return &result, nil
}
