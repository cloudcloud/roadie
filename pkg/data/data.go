// Package data is the general interface to all of the data points across roadie,
// providing the method by which all sources and destinations are interacted with.
package data

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/cloudcloud/roadie/pkg/destinations"
	"github.com/cloudcloud/roadie/pkg/sources"
	"github.com/cloudcloud/roadie/pkg/types"
)

const (
	StateSuccess = "success"
	StateFail    = "fail"
	StateUnknown = "unknown"
)

// Data is the base store for working with all data, containing the configuration
// this instance of roadie was provided with.
type Data struct {
	Content types.Configuration

	l string
	c types.Configer
}

// New will provision the data storage from the provided configuration.
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

// AddHistory will accept some detail about a particular event that has occured.
func (d *Data) AddHistory(s types.Source, t types.Destination, e string, state string) error {
	h := types.History{
		Destination: t,
		OccurredAt:  time.Now(),
		Pattern:     e,
		Source:      s,
		State:       state,
	}

	d.Content.Histories = append(d.Content.Histories, h)
	return d.Write()
}

// Copy will carry out a copy operation based on the incoming execution request.
func (d *Data) Copy(b types.ExecutePayload) (r types.ExecuteResult) {
	s := d.GetSource(b.SourceName)
	ref := types.Reference{Entry: b.EntryName}
	dest := d.GetDestination(b.DestinationName)

	r.References, r.Error = s.Store.CopyTo(ref, dest)

	if r.Error != nil {
		d.AddHistory(s, dest, b.EntryName, StateFail)
	} else {
		d.AddHistory(s, dest, b.EntryName, StateSuccess)
	}

	return
}

// GetDestination will find and load the requested destination.
func (d *Data) GetDestination(s string) types.Destination {
	for _, x := range d.Content.Destinations {
		if x.Name == s {
			return x
		}
	}

	return types.Destination{}
}

// GetDestinationRefs will find and load the content within the requested destination.
func (d *Data) GetDestinationRefs(s string) []types.Reference {
	return d.GetDestination(s).Store.GetRefs()
}

// GetDestinations will return a list of the available destinations.
func (d *Data) GetDestinations() []types.Destination {
	return d.Content.Destinations
}

// GetHistories will provide a list of the historical record from this configuration.
func (d *Data) GetHistories() []types.History {
	return d.Content.Histories
}

// GetSource will find and load the requested source.
func (d *Data) GetSource(s string) types.Source {
	for _, x := range d.Content.Sources {
		if x.Name == s {
			return x
		}
	}

	return types.Source{}
}

// GetSourceRefs will find and load the content within the requested source.
func (d *Data) GetSourceRefs(s string) []types.Reference {
	return d.GetSource(s).Store.GetRefs()
}

// GetSubSourceRefs will retrieve the sub-details for a specific source.
func (d *Data) GetSubSourceRefs(s, u string) []types.Reference {
	return d.GetSource(s).Store.GetSubRefs(u)
}

// GetSources will provide a list of available sources.
func (d *Data) GetSources() []types.Source {
	return d.Content.Sources
}

// RemoveFile will carry out a removal operation based on the removal request input.
func (d *Data) RemoveFile(b types.RemovePayload) error {
	dest := d.GetDestination(b.DestinationName)
	err := dest.Store.RemoveFile(b.EntryName)

	if err != nil {
		d.AddHistory(types.Source{}, dest, b.EntryName, StateFail)
	} else {
		d.AddHistory(types.Source{}, dest, b.EntryName, StateSuccess)
	}

	return err
}

// Write will take all loaded configuration data and write it back to the provided
// configuration file.
func (d *Data) Write() error {
	j, err := json.MarshalIndent(d.Content, "", "  ")
	if err != nil {
		d.c.GetLogger().With("error_message", err).Error("Couldn't marshal indent.")
		return err
	}

	return ioutil.WriteFile(d.l, j, 0644)
}
