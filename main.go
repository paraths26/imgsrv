package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paraths26/imgsrv/router"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func main() {
	//set logger properties :
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	r := router.NewRouter()
	//set recovery middleware to auto recover in panic
	r.Use(gin.Recovery(), ginlogrus.Logger(log))

	gin.SetMode(gin.ReleaseMode)

	// Start and run the server
	log.Infof("Starting HTTP server ... %v", 80)
	if err := r.Run(":80"); err != nil {
		log.Fatalf("Error starting http server : %v", err)
	}
}
