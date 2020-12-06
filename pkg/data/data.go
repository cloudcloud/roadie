package data

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/cloudcloud/roadie/pkg/types"
)

type Data struct {
	l string
	c types.Configer
	f types.ConfigFile
}

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

	json.Unmarshal(content, &d.f)
	return d
}

func (d *Data) Copy(b types.ExecutePayload) types.ExecuteResult {
	r := types.ExecuteResult{}
	switch b.Source.Type {
	case "s3":
		s3 := NewS3(d.c, d.f, b.Source.Source)
		r = s3.CopyTo(b.Source.Entry, b.Destination)

	case "s3_sync":
	}

	return r
}

func (d *Data) GetDestination(s string) types.Destination {
	for _, x := range d.f.Destinations {
		if x.Name == s {
			return x
		}
	}

	return types.Destination{}
}

func (d *Data) GetDestinationRefs(s string) []types.Reference {
	r := []types.Reference{}
	n := d.GetDestination(s)

	switch n.Type {
	case "local_path":
		m, err := filepath.Glob(n.Location + string(filepath.Separator) + "*")
		if err == nil {
			for _, x := range m {
				r = append(r, types.Reference{Entry: x})
			}
		} else {
			d.c.GetLogger().With("error_message", err, "path", n.Location).Error("Unable to load files.")
		}

	}

	return r
}

func (d *Data) GetDestinations() []types.Destination {
	return d.f.Destinations
}

func (d *Data) GetHistories() []types.History {
	return d.f.Histories
}

func (d *Data) GetSource(s string) types.Source {
	for _, x := range d.f.Sources {
		if x.Name == s {
			return x
		}
	}

	return types.Source{}
}

func (d *Data) GetSourceRefs(s string) []types.Reference {
	r := []types.Reference{}
	n := d.GetSource(s)

	switch n.Type {
	case "local_path_recursive":
		fallthrough
	case "local_path":
		m, err := filepath.Glob(n.Location + string(filepath.Separator) + "*")
		if err == nil {
			for _, x := range m {
				r = append(r, types.Reference{Entry: x})
			}
		} else {
			d.c.GetLogger().With("error_message", err, "path", n.Location).Error("Unable to load files.")
		}

	case "s3":
		s3 := NewS3(d.c, d.f, n)
		r = s3.RetrieveListing()

	case "s3_sync":
		s3 := NewS3Sync(d.c, d.f, n)
		r = s3.RetrieveListing()

	}

	return r
}

func (d *Data) GetSources() []types.Source {
	return d.f.Sources
}
