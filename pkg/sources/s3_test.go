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
)

func TestGetMatchingRefsEmptyResponse(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s3mock := mocks.NewMockS3API(ctrl)
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

	s3mock := mocks.NewMockS3API(ctrl)
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

	s3mock := mocks.NewMockS3API(ctrl)
	s3mock.EXPECT().ListObjectsV2(gomock.Any()).Times(1).Return(&s3.ListObjectsV2Output{}, errors.New("err"))

	assert.NotPanics(func() {
		s := &S3{Bucket: "bucket", Path: "/", s3: s3mock}
		refs := s.GetMatchingRefs(types.Reference{Entry: "file"})

		assert.Equal(0, len(refs), "No results when an error occurs.")
	})
}
