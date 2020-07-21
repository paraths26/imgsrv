package handler

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paraths26/imgsrv/models"
	"github.com/paraths26/imgsrv/service/mongo"
	"github.com/paraths26/imgsrv/service/notification"
	"github.com/paraths26/imgsrv/service/storage"
	log "github.com/sirupsen/logrus"
)

//ImgHandler :
type ImgHandler struct {
	MongoSrv   mongo.Client
	StorageSrv storage.S3Srv
	KafkaSrv   notification.Client
}

//Upload :
func (i *ImgHandler) Upload(c *gin.Context) {
	var req models.ImgReq
	if err := c.BindJSON(&req); err != nil {
		log.Errorf("Error unmarshaling image upload request : %v", err)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	imgData, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		log.Errorf("Error reading  image data : %v", err)
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	fn := models.GetFileName(req.Name, req.Album)
	s3Input := models.MakeS3InObject(imgData, fn)
	_, err = i.StorageSrv.PutObject(s3Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error storing image data to s3")
		return
	}
	_, err = i.MongoSrv.InsertOne(c, models.NewImgMetaData(req.Name, req.Album, fn))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error storing image meta data")
		return
	}
	err = i.KafkaSrv.Produce(models.NewKafkaMessage(req.Name, req.Album, "Created"), nil)
	if err != nil {
		c.JSON(http.StatusOK, "Error Sending Notification to Kafka")
		return
	}
	c.JSON(http.StatusOK, "Done")
	return
}

//Delete :
func (i *ImgHandler) Delete(c *gin.Context) {
	imgName := c.Query("imgid")
	album := c.Query("album")

	fn := models.GetFileName(imgName, album)
	s3Input := models.MakeS3DeleteObj(fn)
	_, err := i.StorageSrv.DeleteObject(s3Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error deleting image data to s3")
		return
	}
	_, err = i.MongoSrv.DeleteOne(c, models.GetS3KeyQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error deleting image meta data")
		return
	}
	err = i.KafkaSrv.Produce(models.NewKafkaMessage(imgName, album, "Deleted"), nil)
	if err != nil {
		c.JSON(http.StatusOK, "Error Sending Notification to Kafka")
		return
	}
	c.JSON(http.StatusOK, "Done")
	return
}

//Album :
func (i *ImgHandler) Album(c *gin.Context) {
	albm := c.Query("albumid")
	results, err := i.MongoSrv.Find(c, models.AlbumQuery(albm))
	if err != nil {
		log.Errorf("error fetching images for album %v, %v", albm, err)
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	if data := models.MakeImgResponse(c, results); data != nil {
		c.JSON(http.StatusOK, data)
	}
	c.JSON(http.StatusNoContent, "No images found for given album")
	return
}

//Image :
func (i *ImgHandler) Image(c *gin.Context) {
	albm := c.Query("albumid")
	imgid := c.Query("imgid")
	results, err := i.MongoSrv.Find(c, models.ImgQuery(imgid, albm))
	if err != nil {
		log.Errorf("error fetching images for album %v, %v", albm, err)
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	if data := models.MakeImgResponse(c, results); data != nil {
		c.JSON(http.StatusOK, data)
	}
	c.JSON(http.StatusNoContent, "No images found for given album")
	return
}
