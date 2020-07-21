package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

//S3Srv :
type S3Srv interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
}

//ConnectS3 :
func ConnectS3(id, secret, token, region string) (*s3.S3, error) {
	creds := credentials.NewStaticCredentials(id, secret, token)
	if _, err := creds.Get(); err != nil {
		log.Errorf("Error creating credentials for s3 : %v", err)
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	return s3.New(session.New(), cfg), nil
}
