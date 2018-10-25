package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/wediin/curator/controller"
	"github.com/wediin/curator/lib/db"
)

func newRouter(c Config) (*gin.Engine, error) {
	r := gin.Default()

	photoStorePath := c.StorePath + "/" + c.PhotoDir
	photoClient, err := db.NewPhotoClient(c.MongoServer, c.MongoDB, c.PhotoMongoCollection)
	if err != nil {
		return nil, err
	}

	ping := &controller.PingController{}
	upload := &controller.UploadController{
		Url:            c.Url,
		PhotoStorePath: photoStorePath,
		PhotoRouter:    c.PhotoRouter,
		PhotoDir:       c.PhotoDir,
		PhotoClient:    photoClient,
	}
	graphql := &controller.GraphqlController{
		PhotoClient: photoClient,
	}

	r.Static(c.PhotoRouter, photoStorePath)
	r.GET("/ping", ping.GetController)
	r.POST("/upload", upload.PostController)
	r.GET("/graphql", gin.WrapF(graphql.NewGraphiQLHandlerFunc()))
	r.POST("/graphql", gin.WrapH(graphql.NewHandler()))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	return r, nil
}
