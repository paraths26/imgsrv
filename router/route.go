package router

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/paraths26/imgsrv/config"
)

//NewRouter :
func NewRouter() *gin.Engine {
	r := gin.New()
	api := r.Group("/api")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})
		api.PUT("/image", config.SrvHandler.Upload)
		api.DELETE("/image/:album/:imgid", config.SrvHandler.Delete)
		api.GET("/image/album/:albumid", config.SrvHandler.Album)
		api.GET("/image/album/:albumid/:imgid", config.SrvHandler.Image)
	}

	//debug API end points to be used from go pprof tool
	debug := r.Group("/debug")
	{
		debug.GET("/pprof/heap", gin.WrapH(pprof.Handler("heap")))
		debug.GET("/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))
	}
	return r
}
