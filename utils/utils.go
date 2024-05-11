package utils

import (
	"fmt"
	"net/url"
)

// BuildURL constructs a URL for the Philips Hue API endpoint.
func BuildURL(baseURL, path string, queryParams map[string]string) string {
	u, err := url.Parse(baseURL + path)
	if err != nil {
		// Handle error
		return ""
	}

	// Add query parameters to the URL
	query := u.Query()
	for key, value := range queryParams {
		query.Set(key, value)
	}
	u.RawQuery = query.Encode()

	return u.String()
}

// LogError logs an error message to the console.
func LogError(err error) {
	fmt.Printf("Error: %s\n", err)
}

// Other utility functions...
