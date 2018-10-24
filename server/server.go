package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Init(c Config) error {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r, err := SetupRouter(c)
	if err != nil {
		return err
	}

	port := strconv.Itoa(c.Port)
	r.Run(":" + port)
	return nil
}
