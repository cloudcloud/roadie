package sources

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cloudcloud/roadie/pkg/types"
)

// S3 is a data object that allows working with an S3 location.
type S3 struct {
	Bucket  string `json:"bucket"`
	Depth   int    `json:"depth,omitempty"`
	Path    string `json:"path"`
	Profile string `json:"profile,omitempty"`

	c          types.Configer
	downloader *s3manager.Downloader
	s3         *s3.S3
}

// NewS3
func NewS3(c types.Configer) *S3 {
	e := session.Must(session.NewSession())

	return &S3{
		c:          c,
		downloader: s3manager.NewDownloader(e),
		s3:         s3.New(e),
	}
}

// CopyTo
func (s *S3) CopyTo(r types.Reference, d types.Destination) (list []types.Reference, err error) {
	refs := s.GetMatchingRefs(r)
	switch d.Type {
	case "local_path":
		s.copyToPath(refs, d.Store.GetLocation())
	}

	return
}

func (s *S3) copyToPath(refs []types.Reference, p string) {
	for _, x := range refs {
		filename := strings.TrimPrefix(x.Entry, s.Path)

		f, err := os.Create(p + string(os.PathSeparator) + filename)
		if err != nil {
			s.c.GetLogger().With("error_message", err).Error("Unable to setup local file.")
			continue
		}

		n, err := s.downloader.Download(f, &s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(s.Path + filename),
		})
		if err != nil {
			s.c.GetLogger().With("error_message", err).Error("Could not download from S3.")
		} else {
			s.c.GetLogger().With("downloaded", n, "filename", filename, "bucket", s.Bucket).Info("Completed download from S3.")
		}
	}
}

// GetMatchingRefs
func (s *S3) GetMatchingRefs(r types.Reference) (refs []types.Reference) {
	prefix := s.Path + "/" + r.Entry

	in := &s3.ListObjectsV2Input{
		Bucket:  aws.String(s.Bucket),
		MaxKeys: aws.Int64(1000),
		Prefix:  aws.String(prefix),
	}

	res, err := s.s3.ListObjectsV2(in)
	if err != nil {
		log.Printf("Had this error, [%s]\n", err)
	}

	for _, x := range res.Contents {
		refs = append(refs, types.Reference{Entry: *x.Key})
	}

	return
}

// GetRefs
func (s *S3) GetRefs() (r []types.Reference) {
	delim := "."
	prefix := s.Path + "/"
	if s.Depth > 0 {
		delim = "/"
	}

	in := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.Bucket),
		Delimiter: aws.String(delim),
		MaxKeys:   aws.Int64(1000),
		Prefix:    aws.String(prefix),
	}

	res, err := s.s3.ListObjectsV2(in)
	if err != nil {
		log.Printf("Had an error [%s]\n", err)
	}

	for _, x := range res.CommonPrefixes {
		r = append(r, types.Reference{Entry: strings.TrimRight(strings.TrimLeft(*x.Prefix, prefix), delim)})
	}

	return
}

// Type provides the string reference for this type of source.
func (s *S3) Type() string {
	return "s3"
}
