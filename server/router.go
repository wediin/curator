package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wediin/curator/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)
	upload := new(controller.UploadController)
	graphql := new(controller.GraphqlController)

	r.GET("/ping", ping.GetController)
	r.GET("/graphql", gin.WrapF(graphql.NewGraphiQLHandlerFunc()))
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	r.POST("/upload", upload.PostController)
	return r
}
