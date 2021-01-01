package sources

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudcloud/roadie/pkg/types"
)

// LocalPath
type LocalPath struct {
	Location string `json:"location"`

	c types.Configer
}

// NewLocalPath
func NewLocalPath(c types.Configer) *LocalPath {
	return &LocalPath{c: c}
}

// CopyTo will accept a Reference and Destination to copy the
// refered file to the provided destination.
func (l *LocalPath) CopyTo(r types.Reference, d types.Destination) (list []types.Reference, err error) {
	switch d.Type {
	case "local_path":
		err = l.copyToPath(r, d.Store.GetLocation())
		if err == nil {
			list = append(list, r)
		}
	}

	return
}

func (l *LocalPath) copyToPath(r types.Reference, d string) error {
	in, err := os.Open(r.Entry)
	if err != nil {
		l.c.GetLogger().With("error_message", err).Error("Unable to open file for copying.")
		return err
	}

	out, err := os.Create(d + string(filepath.Separator) + strings.TrimPrefix(r.Entry, l.Location))
	if err != nil {
		l.c.GetLogger().With("error_message", err).Error("Unable to create copy file.")
		return err
	}

	_, err = io.Copy(out, in)
	return err
}

// GetRefs
func (l *LocalPath) GetRefs() (r []types.Reference) {
	m, err := filepath.Glob(l.Location + string(filepath.Separator) + "*")
	if err == nil {
		for _, x := range m {
			r = append(r, types.Reference{Entry: x})
		}
	} else {
		l.c.GetLogger().With("error_message", err, "path", l.Location).Error("Unable to load files.")
	}

	return
}

// Type
func (l *LocalPath) Type() string {
	return "local_path"
}
