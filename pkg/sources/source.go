// Package sources provides implementations for the different types of source
// configurations available.
package sources

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cloudcloud/roadie/pkg/types"
)

// New will provision a new source instance based on the requested type.
func New(t string, c types.Configer) types.Sourcer {
	switch t {
	case "s3":
		return NewS3(c)
	case "local_path":
		return NewLocalPath(c)
	}

	return nil
}

// CreateNew will take a submission and prepare for utilisation.
func CreateNew(s types.ConfigAddSource) types.Source {
	n := types.Source{
		Href: sourceURL(s.Name),
		Name: s.Name,
		Type: s.Type,
	}

	switch s.Type {
	case "local_path":
		n.Store = &LocalPath{Location: s.Path}
		blob, _ := json.Marshal(n.Store)
		n.Config = blob

	case "s3":
		n.Store = &S3{Bucket: s.Bucket, Path: s.Path, Depth: s.Depth}
		blob, _ := json.Marshal(n.Store)
		n.Config = blob

	}

	return n
}

// UpdateExisting will take a new and existing source, generating a new one that
// combines the editable details with the existing others.
func UpdateExisting(e types.ConfigAddSource, s types.Source) types.Source {
	n := types.Source{
		Href: sourceURL(s.Name),
		Name: s.Name,
		Type: s.Type,
	}

	switch s.Type {
	case "local_path":
		n.Store = &LocalPath{Location: e.Path}
		blob, _ := json.Marshal(n.Store)
		n.Config = blob

	case "s3":
		n.Store = &S3{Bucket: e.Bucket, Path: e.Path, Depth: e.Depth}
		blob, _ := json.Marshal(n.Store)
		n.Config = blob

	}

	return n
}

// FromURL will take an expected source name from an external location
// and turn it into what can be assumed as the internal reference.
func FromURL(n string) string {
	x, _ := url.PathUnescape(n)
	return x
}

// PrepareList will decorate a list of sources with additional details
// that are mostly useful externally.
func PrepareList(s []types.Source) (o []types.Source) {
	o = make([]types.Source, len(s))

	for a, x := range s {
		o[a] = PrepareSource(x)
	}

	return
}

// PrepareSource will decorate a single Source entry with external details.
func PrepareSource(s types.Source) (o types.Source) {
	o.Href = sourceURL(s.Name)
	o.Name = s.Name
	o.Type = s.Type

	b, _ := json.Marshal(s.Store)
	o.Config = json.RawMessage(b)

	return
}

func sourceURL(l string) string {
	return fmt.Sprintf("/sources/%s", url.PathEscape(l))
}
