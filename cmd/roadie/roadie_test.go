package main

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	v := m.Run()

	//if v == 0 {
	//os.Exit(1)
	//}
	os.Exit(v)
}

func TestMainPanicConfig(t *testing.T) {
	assert := assert.New(t)

	monkey.Patch(os.Exit, func(i int) {
		panic("os.Exit called")
	})
	defer monkey.UnpatchAll()

	assert.Panics(func() {
		os.Setenv("CONFIG_FILE", "bogus_file")
		os.Setenv("HOSTNAME", "...")
		os.Setenv("PORT", "1")

		main()
	})
}
