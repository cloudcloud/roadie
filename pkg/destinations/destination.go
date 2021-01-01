package destinations

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cloudcloud/roadie/pkg/types"
)

// New
func New(t string, c types.Configer) types.Destinationer {
	switch t {
	case "local_path":
		return NewLocalPath(c)
	}

	return nil
}

// FromURL
func FromURL(u string) string {
	x, _ := url.PathUnescape(u)
	return x
}

// PrepareList
func PrepareList(d []types.Destination) (o []types.Destination) {
	for _, x := range d {
		t := types.Destination{
			Href: destinationURL(x.Name),
			Name: x.Name,
			Type: x.Type,
		}

		b, _ := json.Marshal(x.Store)
		t.Config = json.RawMessage(b)

		o = append(o, t)
	}

	return
}

func destinationURL(l string) string {
	return fmt.Sprintf("/destinations/%s", url.PathEscape(l))
}
