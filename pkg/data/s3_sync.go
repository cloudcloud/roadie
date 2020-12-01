package data

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudcloud/roadie/pkg/types"
)

type S3Sync struct {
	Source types.Source `json:"source"`

	c types.Configer
	f types.ConfigFile
	s *s3.S3
}

func NewS3Sync(c types.Configer, f types.ConfigFile, s types.Source) S3Sync {
	return S3Sync{
		Source: s,
		c:      c,
		f:      f,
		s:      s3.New(session.New()),
	}
}

func (s *S3Sync) RetrieveListing() []types.Reference {
	r := []types.Reference{}
	in := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.Source.Bucket),
		Delimiter: aws.String("/"),
		MaxKeys:   aws.Int64(10000),
		Prefix:    aws.String(s.Source.Path),
	}

	res, err := s.s.ListObjectsV2(in)
	if err != nil {
		s.c.GetLogger().With("error_message", err).Error("Unable to communicate with S3")
	}

	for _, x := range res.CommonPrefixes {
		r = append(r, types.Reference{Entry: strings.TrimLeft(strings.TrimRight(*x.Prefix, "/"), s.Source.Path)})
	}

	return r
}
