package sources

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
		{
			IsNil: false,
			Type:  "s3",
		},
	}

	for _, x := range cases {
		actual := New(x.Type, &config.Config{})

		if x.IsNil {
			assert.Nil(actual, "An unknown type should not provision a Source.")
		} else {
			assert.NotNil(actual, "A valid type provides the Source.")
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

func TestSourceURL(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "pathway",
			Expected: "/sources/pathway",
		},
		{
			Input:    "Hello world",
			Expected: "/sources/Hello%20world",
		},
	}

	for _, x := range cases {
		actual := sourceURL(x.Input)

		assert.Equal(x.Expected, actual, "The generated source URL should match.")
	}
}

func TestPrepareList(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Sources  []types.Source
		Expected []types.Source
		Message  string
	}{
		{
			Sources:  []types.Source{},
			Expected: []types.Source{},
			Message:  "An empty list of Sources is empty.",
		},
		{
			Sources:  []types.Source{types.Source{}},
			Expected: []types.Source{types.Source{Href: "/sources/", Config: json.RawMessage([]byte(`null`))}},
			Message:  "Default single Source in the list is processed as expected.",
		},
	}

	for _, x := range cases {
		actual := PrepareList(x.Sources)

		assert.Len(actual, len(x.Expected), "Same numner of elements provided by PrepareList.")
		assert.Equal(x.Expected, actual, x.Message)
	}
}
