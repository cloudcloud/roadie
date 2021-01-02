package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Port     string
		Error    bool
		Expected *Config
	}{
		{
			Error:    false,
			Expected: &Config{Listener: ":8008", Location: "/config.roadie.json"},
			Port:     "",
		},
	}

	for _, x := range cases {
		os.Setenv("PORT", x.Port)

		actual, err := New()
		assert.Equal(x.Expected.Listener, actual.GetListener())
		assert.Equal(x.Expected.Location, actual.GetConfigFile())

		if !x.Error {
			assert.Nil(err)
		} else {
			assert.NotNil(err)
		}
	}
}

func TestNopBreakCircleCache(t *testing.T) {
	assert := assert.New(t)

	assert.True(true)
}
