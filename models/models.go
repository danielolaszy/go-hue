package models

// Light represents a Philips Hue light.
type Light struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	State    State  `json:"state"`
	Metadata struct {
		Name      string `json:"name"`
		Archetype string `json:"archetype"`
		Function  string `json:"function"`
	}
	On struct {
		On bool `json:"on"`
	}
	Dimming struct {
		Brightness  float32 `json:"brightness"`
		MinDimLevel float32 `json:"min_dim_level"`
	}
	ColorTemperature struct {
		Mirek       int16   `json:"mirek"`
		MinDimLevel float32 `json:"min_dim_level"`
	}
	// Add other fields as needed
}

// State represents the state of a Philips Hue light.
type State struct {
	On         bool `json:"on"`
	Brightness int  `json:"brightness"`
	Hue        int  `json:"hue"`
	Saturation int  `json:"saturation"`
	// Add other fields as needed
}

// ErrorResponse represents an error response from the Philips Hue API.
type ErrorResponse struct {
	Error struct {
		Type        int    `json:"type"`
		Description string `json:"description"`
	} `json:"error"`
}
