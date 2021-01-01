// Package data
package data

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cloudcloud/roadie/pkg/destinations"
	"github.com/cloudcloud/roadie/pkg/sources"
	"github.com/cloudcloud/roadie/pkg/types"
)

// Data
type Data struct {
	Content types.Configuration

	l string
	c types.Configer
	f types.ConfigFile
}

// New
func New(c types.Configer) *Data {
	d := &Data{
		c: c,
		l: c.GetConfigFile(),
	}

	// load the config file and unmarshal into d
	content, err := ioutil.ReadFile(d.l)
	if err != nil {
		c.GetLogger().With("error_message", err).Fatal("Unable to open config.")
	}

	json.Unmarshal(content, &d.Content)

	// fill in each of the sources
	for a, x := range d.Content.Sources {
		store := sources.New(x.Type, c)

		json.Unmarshal(x.Config, &store)
		d.Content.Sources[a].Store = store
	}

	// fill in each of the destinations
	for a, x := range d.Content.Destinations {
		dest := destinations.New(x.Type, c)

		json.Unmarshal(x.Config, &dest)
		d.Content.Destinations[a].Store = dest
	}

	return d
}

// Copy
func (d *Data) Copy(b types.ExecutePayload) (r types.ExecuteResult) {
	s := d.GetSource(b.SourceName)
	ref := types.Reference{Entry: b.EntryName}
	dest := d.GetDestination(b.DestinationName)

	r.References, r.Error = s.Store.CopyTo(ref, dest)
	return
}

// GetDestination
func (d *Data) GetDestination(s string) types.Destination {
	for _, x := range d.Content.Destinations {
		if x.Name == s {
			return x
		}
	}

	return types.Destination{}
}

// GetDestinationRefs
func (d *Data) GetDestinationRefs(s string) []types.Reference {
	return d.GetDestination(s).Store.GetRefs()
}

// GetDestinations
func (d *Data) GetDestinations() []types.Destination {
	return d.Content.Destinations
}

// GetHistories
func (d *Data) GetHistories() []types.History {
	return d.Content.Histories
}

// GetSource
func (d *Data) GetSource(s string) types.Source {
	for _, x := range d.Content.Sources {
		if x.Name == s {
			return x
		}
	}

	return types.Source{}
}

// GetSourceRefs
func (d *Data) GetSourceRefs(s string) []types.Reference {
	return d.GetSource(s).Store.GetRefs()
}

// GetSources
func (d *Data) GetSources() []types.Source {
	return d.Content.Sources
}

// RemoveFile
func (d *Data) RemoveFile(b types.RemovePayload) interface{} {
	return d.GetDestination(b.DestinationName).Store.RemoveFile(b.EntryName)
}
