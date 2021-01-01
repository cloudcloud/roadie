package destinations

import (
	"os"
	"path/filepath"

	"github.com/cloudcloud/roadie/pkg/types"
)

// LocalPath is the implementation of a local filesystem destination.
type LocalPath struct {
	Location string `json:"location"`

	c types.Configer
}

// NewLocalPath provisions a fresh instance of LocalPath.
func NewLocalPath(c types.Configer) *LocalPath {
	return &LocalPath{c: c}
}

// GetLocation returns the full path for this destination.
func (l *LocalPath) GetLocation() string {
	return l.Location
}

// GetRefs will trawl the filesystem location for the destination and
// generate a list of all files located within it.
func (l *LocalPath) GetRefs() (r []types.Reference) {
	m, err := filepath.Glob(l.Location + string(filepath.Separator) + "*")
	if err == nil {
		for _, x := range m {
			r = append(r, types.Reference{Entry: x})
		}
	}

	return
}

// RemoveFile will take the provided reference and remove it.
func (l *LocalPath) RemoveFile(r string) (err error) {
	l.c.GetLogger().With("reference", r).Info("Removing the file.")

	err = os.Remove(r)
	return
}

// Type provides the referencial name for this destination type.
func (l *LocalPath) Type() string {
	return "local_path"
}
