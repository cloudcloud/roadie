// Package types contains data structures and behaviours that help
// shape the way roadie works with data and how the internal APIs
// can be utilised.
package types

import (
	"encoding/json"
	"time"
)

// Configuration defines the overall structure of the content from
// the full roadie configuration.
type Configuration struct {
	Destinations []Destination `json:"destinations"`
	Domains      []string      `json:"domains,omitempty"`
	Histories    []History     `json:"histories"`
	Sources      []Source      `json:"sources"`
}

// Destination contains the details for each individual possible
// destination with the typed attributes included.
type Destination struct {
	Config json.RawMessage `json:"config"`
	Href   string          `json:"href,omitempty"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`

	Store Destinationer `json:"-"`
}

// History is a single record for a copy execution, or some other
// state modification by roadie.
type History struct {
	Config      json.RawMessage `json:"config,omitempty"`
	Destination Destination     `json:"destination"`
	OccurredAt  time.Time       `json:"occurred_at"`
	Pattern     string          `json:"pattern"`
	Source      Source          `json:"source"`
	State       string          `json:"state"`
}

// Source is a structured representation of the original location
// available for data to be copied from.
type Source struct {
	Config json.RawMessage `json:"config"`
	Href   string          `json:"href,omitempty"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`

	// Store is the unmarshal'd Config data in the Type struct.
	Store Sourcer `json:"-"`
}
