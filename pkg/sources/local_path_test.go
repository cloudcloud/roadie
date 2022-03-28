package sources

import (
	"testing"

	"github.com/cloudcloud/roadie/pkg/mocks"
	"github.com/cloudcloud/roadie/pkg/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCopyToEmpty(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := mocks.NewMockConfiger(ctrl)

	ref := types.Reference{SubPath: "", Entry: ""}
	dest := types.Destination{Type: ""}

	assert.NotPanics(func() {
		l := &LocalPath{c: conf}

		res, err := l.CopyTo(ref, dest)

		assert.Equal(0, len(res), "No results from no input.")
		assert.Nil(err, "No error when nothing processed.")
	})
}

func TestCopyToBadInput(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	z, _ := zap.NewProduction(zap.AddCaller())
	conf := mocks.NewMockConfiger(ctrl)
	conf.EXPECT().GetLogger().Times(1).Return(z.Sugar())

	d := mocks.NewMockDestinationer(ctrl)
	d.EXPECT().GetLocation().Times(1).Return("")

	ref := types.Reference{SubPath: "", Entry: "./bad_input/file"}
	dest := types.Destination{Type: "local_path", Store: d}

	assert.NotPanics(func() {
		l := &LocalPath{c: conf}

		res, err := l.CopyTo(ref, dest)

		assert.Equal(0, len(res), "No results as input was bad.")
		assert.NotNil(err, "An error generated from bad input.")
	})
}

func TestCopyToBadOutput(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	z, _ := zap.NewProduction(zap.AddCaller())
	conf := mocks.NewMockConfiger(ctrl)
	conf.EXPECT().GetLogger().Times(1).Return(z.Sugar())

	d := mocks.NewMockDestinationer(ctrl)
	d.EXPECT().GetLocation().Times(1).Return("./an-obviously/bad_path/pls")

	ref := types.Reference{SubPath: "", Entry: "../../test_data/good_config.json"}
	dest := types.Destination{Type: "local_path", Store: d}

	assert.NotPanics(func() {
		l := &LocalPath{c: conf}

		res, err := l.CopyTo(ref, dest)

		assert.Equal(0, len(res), "No results as output was bad.")
		assert.NotNil(err, "An error should've been generated with bad output.")
	})
}

func TestCopyToGood(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := mocks.NewMockConfiger(ctrl)

	d := mocks.NewMockDestinationer(ctrl)
	d.EXPECT().GetLocation().Times(1).Return("../../test_data")

	ref := types.Reference{SubPath: "", Entry: "s3.go"}
	dest := types.Destination{Type: "local_path", Store: d}

	assert.NotPanics(func() {
		l := &LocalPath{c: conf}

		res, err := l.CopyTo(ref, dest)

		assert.Equal(1, len(res), "A result should be returned from a successful copy.")
		assert.Nil(err, "No error when things go well.")
	})
}

func TestGetLocation(t *testing.T) {
	assert := assert.New(t)

	in := ""
	l := &LocalPath{Location: in}
	assert.Equal(in, l.GetLocation(), "Input matches output.")

	in = "purple/monkey/dishwasher"
	l = &LocalPath{Location: in}
	assert.Equal(in, l.GetLocation(), "Input should match output.")
}

func TestType(t *testing.T) {
	assert := assert.New(t)

	l := &LocalPath{}
	assert.Equal(SourceLocalPath, l.Type(), "LocalPath always returns the Type.")
}

func TestGetRefsBadPath(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	z, _ := zap.NewProduction(zap.AddCaller())
	conf := mocks.NewMockConfiger(ctrl)
	conf.EXPECT().GetLogger().Times(1).Return(z.Sugar())

	l := &LocalPath{Location: "...\\/empty/bad/directory", c: conf}

	assert.NotPanics(func() {
		found := l.GetRefs()

		assert.Equal(0, len(found), "No references found in an unknown location.")
	})
}

func TestGetRefs(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conf := mocks.NewMockConfiger(ctrl)

	l := &LocalPath{Location: "", c: conf}

	assert.NotPanics(func() {
		found := l.GetRefs()

		assert.NotEqual(0, len(found), "A non-zero number of files should be found.")
	})
}
