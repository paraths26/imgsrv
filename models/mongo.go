package models

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//ImgMetaData :
type ImgMetaData struct {
	Image string
	Album string
	S3Key string
}

//NewImgMetaData :
func NewImgMetaData(name, album, fn string) ImgMetaData {
	return ImgMetaData{Image: name, Album: album, S3Key: fn}
}

//GetS3KeyQuery :
func GetS3KeyQuery(fn string) bson.M {
	return bson.M{"key": fn}
}

//AlbumQuery :
func AlbumQuery(album string) bson.M {
	return bson.M{"album": album}
}

//ImgQuery :
func ImgQuery(img, album string) bson.M {
	return bson.M{"image": img, "album": album}
}

//MakeImgResponse :
func MakeImgResponse(ctx context.Context, cur *mongo.Cursor) []*ImgMetaData {
	var data []*ImgMetaData
	for cur.Next(ctx) {
		var imgData ImgMetaData
		if err := cur.Decode(&imgData); err != nil {
			log.Errorf("Error decoding mongo response : %v", err)
			return nil
		}
		data = append(data, &imgData)
	}
	return data
}
