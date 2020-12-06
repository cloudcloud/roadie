package types

import (
	"go.uber.org/zap"
	"time"
)

// Configer is an expected behaviour to be provided by some form
// of configuration struct.
type Configer interface {
	GetConfigFile() string
	GetHostname() string
	GetListener() string
	GetLogger() *zap.SugaredLogger
}

// Destination is a representation of a destination for data to be placed.
type Destination struct {
	Href     string `json:"href"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

// History is a representation of a single execution of the file movement.
type History struct {
	Date        time.Time   `json:"date"`
	Destination Destination `json:"destination"`
	Pattern     string      `json:"pattern"`
	Source      Source      `json:"source"`
}

// Reference is an individual entry that is located within a Source
// or Destination as a candidate for transferring.
type Reference struct {
	Entry string `json:"entry"`
}

// Source is a representation of a  source for data to be pulled from.
type Source struct {
	Bucket   string `json:"bucket"`
	Href     string `json:"href"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"`
}

// SourceReference is a single reference entry located within a source.
type SourceReference struct {
	Reference
	Source
}

// ConfigFile contains the data form of our the configuration for
// this instance of roadie.
type ConfigFile struct {
	// Destinations is the list of destinations that are available to place
	// data from the upstreams.
	Destinations []Destination `json:"destinations"`

	// Domains is a list of acceptable domains to be referenced from a
	// CORS perspective.
	Domains []string `json:"domains"`

	// Histories contains
	Histories []History `json:"histories"`

	// Sources is the list of sources from which data can be selected.
	Sources []Source `json:"sources"`
}

// ExecutePayload defines the incoming data request for carrying out
// a copy at the request of the user.
type ExecutePayload struct {
	// Source is the reference for the originating reference.
	Source SourceReference `json:"source"`

	// Destination is the location for which the source should be
	// copied over to.
	Destination Destination `json:"destination"`
}

// ExecuteResult provides structure on responding to a request for the
// copy execution.
type ExecuteResult struct {
	Source
	Destination
}
