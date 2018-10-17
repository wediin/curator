package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wediin/curator/controller"
)

func SetupRouter(c Config) *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)
	upload := &controller.UploadController{
		c.StorePath,
	}
	graphql := new(controller.GraphqlController)

	r.Static("/usercontent", "./uploadFolder")
	r.GET("/ping", ping.GetController)
	r.POST("/upload", upload.PostController)
	r.GET("/graphql", gin.WrapF(graphql.NewGraphiQLHandlerFunc()))
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	return r
}
