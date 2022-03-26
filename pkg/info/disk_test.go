package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskDetailsEmpty(t *testing.T) {
	assert := assert.New(t)

	input := []string{}
	assert.NotPanics(func() {
		response := DiskDetails(input)

		assert.Empty(response, "There should be no response from the function call.")
	})
}

func TestDiskDetailsCWD(t *testing.T) {
	assert := assert.New(t)

	input := []string{"."}
	assert.NotPanics(func() {
		response := DiskDetails(input)

		assert.Equal(len(input), len(response), "The input and output length should be the same.")
		assert.NotEqual(uint64(0), response[0].Used, "There should be a non-zero amount of Used space.")
		assert.Equal(input[0], response[0].Path, "The input path should be in the response result.")
	})
}
