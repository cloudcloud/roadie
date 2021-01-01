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

// FromURL will take an expected source name from an external location
// and turn it into what can be assumed as the internal reference.
func FromURL(n string) string {
	x, _ := url.PathUnescape(n)
	return x
}

// PrepareList will decorate a list of sources with additional details
// that are mostly useful externally.
func PrepareList(s []types.Source) (o []types.Source) {
	for _, x := range s {
		t := types.Source{
			Href: sourceURL(x.Name),
			Name: x.Name,
			Type: x.Type,
		}

		b, _ := json.Marshal(x.Store)
		t.Config = json.RawMessage(b)

		o = append(o, t)
	}

	return
}

func sourceURL(l string) string {
	return fmt.Sprintf("/sources/%s", url.PathEscape(l))
}
