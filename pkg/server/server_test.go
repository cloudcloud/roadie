package server

import (
	"testing"

	"github.com/cloudcloud/roadie/pkg/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewEmptyConfiger(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	z, _ := zap.NewProduction(zap.AddCaller())
	log := z.Sugar()

	c := mocks.NewMockConfiger(ctrl)
	c.EXPECT().GetConfigFile().Times(1).Return("")
	c.EXPECT().GetLogger().Times(1).Return(log)

	assert.Panics(func() {
		result := New(c)
		assert.Nil(result, "Valid Server should not be provided without a config file.")
	}, "Failure to provide a Config File should panic.")
}

func TestNewBasicConfiger(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	z, _ := zap.NewProduction(zap.AddCaller())
	log := z.Sugar()

	c := mocks.NewMockConfiger(ctrl)
	c.EXPECT().GetConfigFile().Times(1).Return("../../test_data/good_config.json")
	c.EXPECT().GetLogger().Times(1).Return(log)

	assert.NotPanics(func() {
		result := New(c)
		assert.NotNil(result, "Valid server should be provided by New.")
	}, "With a valid Config File a Server should be created.")
}
