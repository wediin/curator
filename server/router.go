package server

import (
	"github.com/gin-gonic/gin"
	"github.com/wediin/curator/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)
	graphql := new(controller.GraphqlController)
	graphqlController := gin.WrapH(graphql.NewHandler())

	r.GET("/ping", ping.GetController)
	r.GET("/graphql", graphqlController)
	r.POST("/graphql", graphqlController)

	return r
}
