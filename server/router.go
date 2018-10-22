package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wediin/curator/controller"
)

func SetupRouter(c Config) *gin.Engine {
	r := gin.Default()

	photoStorePath := c.StorePath + "/" + c.PhotoDir

	ping := new(controller.PingController)
	upload := &controller.UploadController{
		Url:                  c.Url,
		MongoServer:          c.MongoServer,
		MongoDB:              c.MongoDB,
		PhotoMongoCollection: c.PhotoMongoCollection,
		PhotoStorePath:       photoStorePath,
		PhotoRouter:          c.PhotoRouter,
		PhotoDir:             c.PhotoDir,
	}
	graphql := &controller.GraphqlController{
		MongoServer:          c.MongoServer,
		MongoDB:              c.MongoDB,
		PhotoMongoCollection: c.PhotoMongoCollection,
	}

	r.Static(c.PhotoRouter, photoStorePath)
	r.GET("/ping", ping.GetController)
	r.POST("/upload", upload.PostController)
	r.GET("/graphql", gin.WrapF(graphql.NewGraphiQLHandlerFunc()))
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	return r
}
