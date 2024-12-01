// Package data is the general interface to all of the data points across roadie,
// providing the method by which all sources and destinations are interacted with.
package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cloudcloud/roadie/pkg/destinations"
	"github.com/cloudcloud/roadie/pkg/sources"
	"github.com/cloudcloud/roadie/pkg/types"
	humanize "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/disk"
)

const (
	StateSuccess = "success"
	StateFail    = "fail"
	StateUnknown = "unknown"

	StateAdded         = "[added]"
	StateRemoved       = "[removed]"
	StateFailedRemoval = "[not-removed]"
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
		c.GetLogger().With("error_message", err).Error("Unable to open config.")
		panic("Unable to open config file.")
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

// AddDestination will add a new destination to the list of those available.
func (d *Data) AddDestination(n types.Destination) error {
	for _, x := range d.Content.Destinations {
		if x.Name == n.Name {
			return fmt.Errorf("Destination '%s' already exists.", x.Name)
		}
	}

	d.Content.Destinations = append(d.Content.Destinations, n)
	return d.AddHistory(types.Source{Name: StateAdded}, n, n.Name, StateSuccess)
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

// AddSource will add a new source to the list of those available.
func (d *Data) AddSource(n types.Source) error {
	for _, x := range d.Content.Sources {
		if x.Name == n.Name {
			return fmt.Errorf("Source '%s' already exists.", x.Name)
		}
	}

	d.Content.Sources = append(d.Content.Sources, n)
	return d.AddHistory(n, types.Destination{Name: StateAdded}, n.Name, StateSuccess)
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

func (d *Data) GetDestinationsWithDetails() []any {
	output := []any{}

	destinations := d.Content.Destinations
	for _, v := range destinations {
		usage, err := disk.Usage(v.Store.GetLocation())
		if err != nil {
			d.c.GetLogger().With("error_message", err, "path", v.Store.GetLocation()).Error("Unable to determine free disk space.")
		}

		output = append(output, map[string]interface{}{
			"config":    v.Config,
			"href":      v.Href,
			"name":      v.Name,
			"type":      v.Type,
			"disk_free": humanize.Bytes(usage.Free),
		})
	}

	return output
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

// RemoveDestination will take a specific destination name and remove it from our list.
func (d *Data) RemoveDestination(s string) error {
	dest := d.GetDestination(s)
	if dest.Name == s && s != "" {
		tmpDests := []types.Destination{}
		dests := d.Content.Destinations

		for _, x := range dests {
			if x.Name != s {
				tmpDests = append(tmpDests, x)
			}
		}

		d.Content.Destinations = tmpDests

		d.AddHistory(types.Source{}, dest, StateRemoved, StateSuccess)
		return nil
	}

	d.AddHistory(types.Source{}, types.Destination{Name: s}, StateFailedRemoval, StateFail)
	return fmt.Errorf("Could not find the '%s' destination to remove.", s)
}

// RemoveSource will take a specific source name and remove it from our list.
func (d *Data) RemoveSource(s string) error {
	source := d.GetSource(s)
	if source.Name == s && s != "" {
		tmpSources := []types.Source{}
		sources := d.Content.Sources

		for _, x := range sources {
			if x.Name != s {
				tmpSources = append(tmpSources, x)
			}
		}

		d.Content.Sources = tmpSources
		d.AddHistory(source, types.Destination{}, StateRemoved, StateSuccess)
		return nil
	}

	d.AddHistory(types.Source{Name: s}, types.Destination{}, StateFailedRemoval, StateFail)
	return fmt.Errorf("Could not find the '%s' source to remove.", s)
}

// UpdateDestination will take a copy of a Destination and replace it internally
// based on the Name.
func (d *Data) UpdateDestination(b types.Destination) error {
	if b.Name == "" {
		return fmt.Errorf("Invalid destination name '%s' provided.", b.Name)
	}

	for idx, x := range d.Content.Destinations {
		if x.Name == b.Name {
			d.Content.Destinations[idx] = b
			return nil
		}
	}

	return fmt.Errorf("Unable to find destination '%s'.", b.Name)
}

// UpdateSource will take a copy of a Source and replace it internally based on the Name.
func (d *Data) UpdateSource(s types.Source) error {
	if s.Name == "" {
		return fmt.Errorf("Invalid source name '%s' provided.", s.Name)
	}

	for idx, x := range d.Content.Sources {
		if x.Name == s.Name {
			d.Content.Sources[idx] = s
			return nil
		}
	}

	return fmt.Errorf("Unable to find source '%s'.", s.Name)
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
