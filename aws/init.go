package aws

import (
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	kinesissvc *kinesis.Kinesis
	s3svc      *s3.S3
	region     string
	profile    string
)

func init() {
	// AWS Related
	flag.StringVar(&region, "region", "us-east-1", "Region to be used for AWS")
	flag.StringVar(&profile, "profile", "default", "Profile to be used for AWS")

}

func InitSvc() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: profile,
	}))

	if kinesissvc == nil {
		kinesissvc = kinesis.New(sess, &aws.Config{
			Region: aws.String(region),
		})
	}
	if s3svc == nil {
		s3svc = s3.New(sess, &aws.Config{
			Region: aws.String(region),
		})
	}
}
