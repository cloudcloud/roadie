package destinations

import (
	"encoding/json"
	"testing"

	"github.com/cloudcloud/roadie/pkg/config"
	"github.com/cloudcloud/roadie/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		IsNil bool
		Type  string
	}{
		{
			IsNil: true,
			Type:  "invalid_type",
		},
		{
			IsNil: false,
			Type:  "local_path",
		},
	}

	for _, x := range cases {
		actual := New(x.Type, &config.Config{})

		if x.IsNil {
			assert.Nil(actual, "An unknown type should not provision a Destination.")
		} else {
			assert.NotNil(actual, "A valid type provides the Destination.")
		}
	}
}

func TestFromURL(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "hello%20world",
			Expected: "hello world",
		},
		{
			Input:    "SomethingElse",
			Expected: "SomethingElse",
		},
		{
			Input:    "",
			Expected: "",
		},
	}

	for _, x := range cases {
		actual := FromURL(x.Input)

		assert.Equal(x.Expected, actual, "FromURL should provide the mapped value.")
	}
}

func TestDestinationURL(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "pathway",
			Expected: "/destinations/pathway",
		},
		{
			Input:    "Hello world",
			Expected: "/destinations/Hello%20world",
		},
	}

	for _, x := range cases {
		actual := destinationURL(x.Input)

		assert.Equal(x.Expected, actual, "The generated destination URL should match.")
	}
}

func TestPrepareList(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Destinations []types.Destination
		Expected     []types.Destination
		Message      string
	}{
		{
			Destinations: []types.Destination{},
			Expected:     []types.Destination{},
			Message:      "An empty list of Destinations is empty.",
		},
		{
			Destinations: []types.Destination{types.Destination{}},
			Expected:     []types.Destination{types.Destination{Href: "/destinations/", Config: json.RawMessage([]byte(`null`))}},
			Message:      "Default single Destination in the list is processed as expected.",
		},
	}

	for _, x := range cases {
		actual := PrepareList(x.Destinations)

		assert.Len(actual, len(x.Expected), "Same numner of elements provided by PrepareList.")
		assert.Equal(x.Expected, actual, x.Message)
	}
}
