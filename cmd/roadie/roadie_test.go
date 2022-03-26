package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	v := m.Run()

	os.Exit(v)
}

func TestMainPanicConfig(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		os.Setenv("CONFIG_FILE", "bogus_file")
		os.Setenv("HOSTNAME", "...")
		os.Setenv("PORT", "1")

		main()
	})
}
