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
