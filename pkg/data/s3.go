package data

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cloudcloud/roadie/pkg/types"
)

// S3 is a location type that works with single files within an S3 location.
type S3 struct {
	Source types.Source `json:"source"`

	c types.Configer
	d *s3manager.Downloader
	f types.ConfigFile
	s *s3.S3
}

// NewS3 will provision an S3 location handler.
func NewS3(c types.Configer, f types.ConfigFile, s types.Source) S3 {
	e := session.Must(session.NewSession())
	return S3{
		Source: s,
		c:      c,
		d:      s3manager.NewDownloader(e),
		f:      f,
		s:      s3.New(e),
	}
}

// CopyTo will take another location and carry out a copy of the specified file
// into that second location.
func (s *S3) CopyTo(e string, n types.Destination) types.ExecuteResult {
	r := types.ExecuteResult{Source: s.Source, Destination: n}

	// list all the files matching the path
	entries := s.getMatchingEntries(e)
	s.c.GetLogger().With("entries", entries).Info("Pulled entries for the CopyTo.")

	if len(entries) < 1 {
		return r
	}

	// for each, copy across
	for _, x := range entries {
		filename := strings.TrimPrefix(x.Entry, s.Source.Path)
		switch n.Type {
		case "local_path":
			f, err := os.Create(n.Location + string(os.PathSeparator) + filename)
			if err != nil {
				s.c.GetLogger().With("error_message", err).Error("Unable to setup local file.")
				continue
			}

			n, err := s.d.Download(f, &s3.GetObjectInput{
				Bucket: aws.String(s.Source.Bucket),
				Key:    aws.String(s.Source.Path + filename),
			})
			if err != nil {
				s.c.GetLogger().With("error_message", err).Error("Could not download from S3.")
			} else {
				s.c.GetLogger().With("downloaded", n, "filename", filename, "bucket", s.Source.Bucket).Info("Completed download from S3.")
			}
		}
	}

	return r
}

// RetrieveListing uses the attributes for the S3 location to generate a listing
// of data contained within the location.
func (s *S3) RetrieveListing() []types.Reference {
	r := []types.Reference{}
	in := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.Source.Bucket),
		Delimiter: aws.String("."),
		MaxKeys:   aws.Int64(1000),
		Prefix:    aws.String(s.Source.Path),
	}

	res, err := s.s.ListObjectsV2(in)
	if err != nil {
		s.c.GetLogger().With("error_message", err).Error("Unable to communicate with S3")
	}

	for _, x := range res.CommonPrefixes {
		r = append(r, types.Reference{Entry: strings.TrimRight(strings.TrimLeft(*x.Prefix, s.Source.Path), ".")})
	}

	return r
}

func (s *S3) getMatchingEntries(e string) []types.Reference {
	r := []types.Reference{}
	in := &s3.ListObjectsV2Input{
		Bucket:  aws.String(s.Source.Bucket),
		MaxKeys: aws.Int64(1000),
		Prefix:  aws.String(s.Source.Path + e),
	}

	res, err := s.s.ListObjectsV2(in)
	if err != nil {
		s.c.GetLogger().With("error_message", err).Error("Unable to retrieve the listing from S3")
	}

	for _, x := range res.Contents {
		r = append(r, types.Reference{Entry: *x.Key})
	}

	return r
}
