package data

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/cloudcloud/roadie/pkg/config"
	"github.com/cloudcloud/roadie/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Config types.Configer
	}{
		{
			Config: &config.Config{Location: "../../test_data/good_config.json"},
		},
	}

	for _, x := range cases {
		actual := New(x.Config)

		assert.Equal(x.Config, actual.c)
	}
}

func TestNewNoConfigFile(t *testing.T) {
	assert := assert.New(t)

	monkey.Patch(os.Exit, func(i int) {
		panic("os.Exit called")
	})
	defer monkey.UnpatchAll()

	assert.Panics(func() {
		New(&config.Config{Location: "bogus_file"})
	})
}

func TestGetDestinations(t *testing.T) {
	assert := assert.New(t)

	d := &Data{Content: types.Configuration{}}
	assert.Len(d.GetDestinations(), 0, "Base initialisation should have no Destinations.")

	d.Content.Destinations = append(d.Content.Destinations, types.Destination{})
	assert.Len(d.GetDestinations(), 1, "Adding a single Destination will appear with GetDestinations()")
}

func TestGetHistories(t *testing.T) {
	assert := assert.New(t)

	d := &Data{Content: types.Configuration{}}
	assert.Len(d.GetHistories(), 0, "Base initialisation should have no Histories.")

	d.Content.Histories = append(d.Content.Histories, types.History{})
	assert.Len(d.GetHistories(), 1, "Adding a single History will appear with GetHistories()")
}

func TestGetSources(t *testing.T) {
	assert := assert.New(t)

	d := &Data{Content: types.Configuration{}}
	assert.Len(d.GetSources(), 0, "Base initialisation should have no Sources.")

	d.Content.Sources = append(d.Content.Sources, types.Source{})
	assert.Len(d.GetSources(), 1, "Adding a single Source will appear with GetSources()")
}

func TestAddHistory(t *testing.T) {
	assert := assert.New(t)

	sour := types.Source{}
	dest := types.Destination{}
	patt := "test"
	state := StateSuccess

	l := "test_data/tmp.file"
	c := &config.Config{Location: l}
	d := &Data{c: c, l: l}

	err := d.AddHistory(sour, dest, patt, state)
	assert.Nil(err)

	freshData := New(c)
	assert.Len(freshData.GetHistories(), 1, "A single History should've been added.")

	err = d.AddHistory(sour, dest, patt, state)
	assert.Nil(err)

	freshData = New(c)
	assert.Len(freshData.GetHistories(), 2, "Adding another History makes the length 2.")
	os.RemoveAll(l)
}
