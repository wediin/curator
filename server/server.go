package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Init(c Config) {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := SetupRouter(c)

	port := strconv.Itoa(c.Port)
	r.Run(":" + port)
}
