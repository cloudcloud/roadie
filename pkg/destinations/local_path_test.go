package destinations

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudcloud/roadie/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Location string
	}{
		{
			Location: "",
		},
		{
			Location: "path",
		},
	}

	for _, x := range cases {
		l := &LocalPath{Location: x.Location}
		assert.Equal(x.Location, l.GetLocation(), "Location should match provided value.")
	}
}

func TestGetRefs(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Location string
		Count    int
	}{
		{
			Location: "test_data",
			Count:    1,
		},
		{
			Location: "../destinations",
			Count:    5,
		},
	}

	for _, x := range cases {
		l := &LocalPath{Location: x.Location}

		actual := l.GetRefs()
		assert.GreaterOrEqual(len(actual), x.Count, "Specific number of files should've been found.")
	}
}

func TestRemoveFile(t *testing.T) {
	assert := assert.New(t)

	loc := "test_data"
	input := loc + string(filepath.Separator) + ".gitkeep"
	output := loc + string(filepath.Separator) + "test_file"

	in, err := os.Open(input)
	if err != nil {
		assert.FailNow(err.Error(), "Unable to open input file.")
	}
	out, err := os.Create(output)
	if err != nil {
		assert.FailNow(err.Error(), "Unable to create output file.")
	}

	_, err = io.Copy(out, in)
	if err != nil {
		assert.FailNow(err.Error(), "Unable to copy to test file.")
	}

	l := &LocalPath{Location: loc, c: &config.Config{}}
	err = l.RemoveFile(output)
	assert.Nil(err, "No error should be encountered when removing the known file.")
}

func TestType(t *testing.T) {
	assert := assert.New(t)

	l := &LocalPath{}
	assert.Equal("local_path", l.Type(), "The Type of LocalPath should be local_path.")
}
