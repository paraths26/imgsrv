// package gcp

// import (
// 	"context"

// 	"cloud.google.com/go/storage"
// 	log "github.com/sirupsen/logrus"
// 	"google.golang.org/api/option"
// )

// //ConnectGCPStorage :
// func ConnectGCPStorage(jsonPath, project, bucketName string) (*storage.BucketHandle, error) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
// 	if err != nil {
// 		log.Errorf("Error connecting to GCP storage service : %v", err)
// 		return nil, err
// 	}

// 	return client.Bucket(bucketName).Object("obj").NewWriter(), nil
// 	storage.BucketHandle
// }
