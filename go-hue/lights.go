package gohue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danielolaszy/go-hue/client"
	"github.com/danielolaszy/go-hue/errors"
	"github.com/danielolaszy/go-hue/models"
)

// Lights represents methods for interacting with lights endpoints of the Philips Hue API.
type Lights struct {
	client *client.Client
}

// NewLights creates a new instance of the Lights struct.
func NewLights(c *client.Client) *Lights {
	return &Lights{client: c}
}

// GetLights retrieves information about all lights from the Philips Hue API.
func (l *Lights) GetLights() ([]models.Light, error) {
	resp, err := l.client.MakeRequest(http.MethodGet, "/light", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := new(errors.APIError)
		if err := json.NewDecoder(resp.Body).Decode(apiErr); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %w", err)
		}
		return nil, apiErr
	}

	var lights map[string]models.Light
	if err := json.NewDecoder(resp.Body).Decode(&lights); err != nil {
		return nil, fmt.Errorf("failed to decode lights response: %w", err)
	}

	var lightList []models.Light
	for _, light := range lights {
		lightList = append(lightList, light)
	}

	return lightList, nil
}

// GetLight retrieves information about a specific light by its ID from the Philips Hue API.
// func (l *Lights) GetLight(lightID string) (models.Light, error) {
// 	// Similar implementation to GetLights, but making a request to /lights/<lightID>
// 	_, err := l.client.MakeRequest(http.MethodGet, "/lights", nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return models.Light{}, nil // Placeholder return statement
// }

// SetLightState sets the state of a specific light by its ID using the Philips Hue API.
func (l *Lights) SetLightState(lightID string, state models.State) error {
	// Implementation to send a PUT request to /lights/<lightID>/state with the specified state
	return nil // Placeholder return statement
}

// Other methods for interacting with lights endpoints...
