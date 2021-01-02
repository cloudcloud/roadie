// Package destinations provides implementations for the different types
// of destination.
package destinations

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cloudcloud/roadie/pkg/types"
)

// New will provision an instance of the requested destination type.
func New(t string, c types.Configer) types.Destinationer {
	switch t {
	case "local_path":
		return NewLocalPath(c)
	}

	return nil
}

// FromURL will take an input destination and convert it to the known name.
func FromURL(u string) string {
	x, _ := url.PathUnescape(u)
	return x
}

// PrepareList is a semi-decorator that will add details to a list of
// destinations that is primarily useful for external contexts.
func PrepareList(d []types.Destination) (o []types.Destination) {
	for _, x := range d {
		o = append(o, PrepareDestination(x))
	}

	return
}

// PrepareDestination will decorate a single destination with details useful
// in an external context.
func PrepareDestination(d types.Destination) (o types.Destination) {
	o.Href = destinationURL(d.Name)
	o.Name = d.Name
	o.Type = d.Type

	b, _ := json.Marshal(d.Store)
	o.Config = json.RawMessage(b)

	return
}

func destinationURL(l string) string {
	return fmt.Sprintf("/destinations/%s", url.PathEscape(l))
}
