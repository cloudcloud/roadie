package sources

import (
	"log"
	"os"
	"path/filepath"
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

// NewS3 will provision an S3 source.
func NewS3(c types.Configer) *S3 {
	e := session.Must(session.NewSession())

	return &S3{
		c:          c,
		downloader: s3manager.NewDownloader(e),
		s3:         s3.New(e),
	}
}

// CopyTo will carry out the copy operation from the current bucket
// configuration into the provided destination.
func (s *S3) CopyTo(r types.Reference, d types.Destination) (list []types.Reference, err error) {
	list = make([]types.Reference, 0)

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

		s.c.GetLogger().With("local_filename", p+string(os.PathSeparator)+filename).Info("Ok")
		err := os.MkdirAll(filepath.Dir(filepath.Clean(p+string(os.PathSeparator)+filename)), 0755)
		if err != nil {
			s.c.GetLogger().With("error_message", err).Error("Couldn't prepare directory")
		}

		f, err := os.Create(p + filename)
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

// GetMatchingRefs will pull a list of references from the bucket by using
// the prefix capability of the S3 API for a specific file path.
func (s *S3) GetMatchingRefs(r types.Reference) (refs []types.Reference) {
	refs = make([]types.Reference, 0)
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

// GetRefs will pull a full list of references for the bucket, taking into
// account the nature of folder-structures with the depth.
func (s *S3) GetRefs() (r []types.Reference) {
	r = make([]types.Reference, 0)
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

// GetSubRefs will
func (s *S3) GetSubRefs(sub string) (r []types.Reference) {
	r = make([]types.Reference, 0)
	delim := "/"
	prefix := s.Path + "/" + sub + "/"

	in := &s3.ListObjectsV2Input{
		Bucket:    aws.String(s.Bucket),
		Delimiter: aws.String(delim),
		MaxKeys:   aws.Int64(1000),
		Prefix:    aws.String(prefix),
	}

	res, err := s.s3.ListObjectsV2(in)
	if err != nil {
		s.c.GetLogger().With("error_message", err).Error("Could not GetSubRefs.")
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
