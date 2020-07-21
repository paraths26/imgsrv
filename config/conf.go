package config

import (
	"github.com/paraths26/imgsrv/handler"
	"github.com/paraths26/imgsrv/service/mongo"
	"github.com/paraths26/imgsrv/service/storage"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
)

type mongoConf struct {
	Host       string `json:"url"`
	DB         string `json:"db"`
	Collection string `json:"coll"`
}

type awsConf struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Bucket string `json:"bucket"`
	Token  string `json:"token"`
	Region string `json:"region"`
}

type kafkaConf struct {
	Host  string `json:"host"`
	Topic string `json:"topic"`
}

var (
	SrvHandler handler.ImgHandler
	// kafkaConf    srvConf
	mongoDetails mongoConf
	awsDetail    awsConf
	kafkaDetail  kafkaConf
)

func init() {
	config.Load(file.NewSource(
		file.WithPath("./config.json"),
	))
	config.Get("kafka").Scan(&kafkaDetail)
	config.Get("mongo").Scan(&mongoDetails)
	config.Get("s3").Scan(&awsDetail)
	mongoColl := mongo.ConnnectMongo(mongoDetails.Host, mongoDetails.DB, mongoDetails.Collection)
	if mongoColl == nil {
		log.Fatal("Error Connecting to Mongo DB")
	}
	SrvHandler.MongoSrv = mongoColl
	s3Srv, err := storage.ConnectS3(awsDetail.ID, awsDetail.Secret, awsDetail.Token, awsDetail.Region)
	if err != nil {
		log.Fatal("Error connecting to s3 service")
	}
	SrvHandler.StorageSrv = s3Srv
	return
}
