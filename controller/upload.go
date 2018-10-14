package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (ctr *UploadController) PostController(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	saveFileName := fmt.Sprintf("%d-%s", time.Now().Unix(), handler.Filename)
	out, err := os.Create(saveFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	io.Copy(out, file)

	c.String(http.StatusOK, "upload successful\n")
}
