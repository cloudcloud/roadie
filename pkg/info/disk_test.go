package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskDetailsEmpty(t *testing.T) {
	assert := assert.New(t)

	input := ""
	assert.NotPanics(func() {
		response := DiskDetails(input)

		assert.Empty(response, "There should be no response from the function call.")
	})
}

func TestDiskDetailsCWD(t *testing.T) {
	assert := assert.New(t)

	input := "."
	assert.NotPanics(func() {
		response := DiskDetails(input)

		assert.NotEqual(Disk{}, response, "The output should not be empty.")
		assert.NotEqual(uint64(0), response.Used, "There should be a non-zero amount of Used space.")
		assert.Equal(input, response.Path, "The input path should be in the response result.")
	})
}
