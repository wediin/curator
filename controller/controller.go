package controller

import (
	"github.com/gin-gonic/gin"
)

func statusError(c *gin.Context, status int, err error) {
	c.Error(err)
	c.String(status, err.Error())
}

func statusOK(c *gin.Context, status int, msg string) {
	c.String(status, msg)
}
