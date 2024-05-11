package errors

import "fmt"

// APIError represents an error returned by the Philips Hue API.
type APIError struct {
	Type        int    `json:"type"`
	Description string `json:"description"`
}

// Error returns a string representation of the API error.
func (e *APIError) Error() string {
	return fmt.Sprintf("Hue API error: %s (type: %d)", e.Description, e.Type)
}
