package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Error    bool
		Expected *Config
		File     string
		Host     string
		Port     string
	}{
		{
			Error:    false,
			Expected: &Config{Hostname: "http://localhost:8080", Listener: ":8008", Location: "/config.roadie.json"},
			File:     "",
			Host:     "",
			Port:     "",
		},
		{
			Error:    false,
			Expected: &Config{Hostname: "...", Listener: ":8008", Location: "/config.roadie.json"},
			File:     "/config.roadie.json",
			Host:     "...",
			Port:     "",
		},
		{
			Error:    false,
			Expected: &Config{Hostname: "...", Listener: ":9999", Location: "/roadie.json"},
			File:     "/roadie.json",
			Host:     "...",
			Port:     "9999",
		},
	}

	for _, x := range cases {
		os.Setenv("CONFIG_FILE", x.File)
		os.Setenv("HOSTNAME", x.Host)
		os.Setenv("PORT", x.Port)

		actual, err := New()
		assert.Equal(x.Expected.Hostname, actual.GetHostname())
		assert.Equal(x.Expected.Listener, actual.GetListener())
		assert.Equal(x.Expected.Location, actual.GetConfigFile())

		if !x.Error {
			assert.Nil(err)
		} else {
			assert.NotNil(err)
		}
	}
}

func TestGetLogger(t *testing.T) {
	assert := assert.New(t)

	c := &Config{}
	found := c.GetLogger()

	assert.NotNil(found)
}
