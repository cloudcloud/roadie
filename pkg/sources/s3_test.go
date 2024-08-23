package sources

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudcloud/roadie/pkg/mocks"
	"github.com/cloudcloud/roadie/pkg/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetMatchingRefsEmptyResponse(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3mock := NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{}, nil)

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", s3: s3mock}
		refs := s.GetMatchingRefs(types.Reference{Entry: "file"})

		assert.Equal(0, len(refs), "No results on standard call with nothing found.")
	})
}

func TestGetMatchingRefsBasicResponse(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3mock := NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{
		Contents: []*s3.Object{
			&s3.Object{
				Key: aws.String("key"),
			},
		},
	}, nil)

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", s3: s3mock}
		refs := s.GetMatchingRefs(types.Reference{Entry: "file"})

		assert.Equal(1, len(refs), "Results on standard call.")
	})
}

func TestGetMatchingRefsError(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// As we chain things through, and use our own interface here, the method signature
	// will only allow the struct to be returned. This means we can validate that the
	// first method in the chain is called, with the subsequent calls being out of our
	// control.
	z, _ := zap.NewProduction()
	l := mocks.NewMockLogger(ctrl)
	l.EXPECT().With("error_message", gomock.Any()).Times(1).Return(z.Sugar())

	c := mocks.NewMockConfiger(ctrl)
	c.EXPECT().GetLogger().Times(1).Return(l)

	s3mock := NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{}, errors.New("err"))

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", s3: s3mock, c: c}
		refs := s.GetMatchingRefs(types.Reference{Entry: "file"})

		assert.Equal(0, len(refs), "No results when an error occurs.")
	})
}

func TestGetRefsEmptyResponse(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3mock := NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{
		CommonPrefixes: []*s3.CommonPrefix{},
		IsTruncated:    aws.Bool(false),
	}, nil)

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", s3: s3mock}
		refs := s.GetRefs()

		assert.Equal(0, len(refs), "No refs when there's an empty response.")
	})
}

func TestGetRefsBasicResponse(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3mock := NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{
		CommonPrefixes: []*s3.CommonPrefix{
			&s3.CommonPrefix{
				Prefix: aws.String("prefix/"),
			},
		},
		IsTruncated: aws.Bool(false),
	}, nil)

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", Depth: 1, s3: s3mock}
		refs := s.GetRefs()

		assert.Equal(1, len(refs), "Refs when there's a response.")
		assert.Equal("prefix", refs[0].Entry, "Suffix should be stripped from the Entry.")
	})
}

func TestGetRefsError(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// As we chain things through, and use our own interface here, the method signature
	// will only allow the struct to be returned. This means we can validate that the
	// first method in the chain is called, with the subsequent calls being out of our
	// control.
	z, _ := zap.NewProduction()
	l := mocks.NewMockLogger(ctrl)
	l.EXPECT().With("error_message", gomock.Any()).Times(1).Return(z.Sugar())

	c := mocks.NewMockConfiger(ctrl)
	c.EXPECT().GetLogger().Times(1).Return(l)

	s3mock := NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{}, errors.New("err"))

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", s3: s3mock, c: c}
		refs := s.GetRefs()

		assert.Equal(0, len(refs), "No results when an error occurs.")
	})
}
