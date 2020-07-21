package models

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	//S3Bucket :
	S3Bucket = "imgSrv"
)

//MakeS3InObject :
func MakeS3InObject(data []byte, fn string) *s3.PutObjectInput {
	fileBytes := bytes.NewReader(data)
	return &s3.PutObjectInput{
		Bucket:      aws.String(S3Bucket),
		Key:         aws.String(fn),
		Body:        fileBytes,
		ContentType: aws.String(".jpg"),
	}
}

//MakeS3DeleteObj :
func MakeS3DeleteObj(fn string) *s3.DeleteObjectInput {
	return &s3.DeleteObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(fn),
	}
}
