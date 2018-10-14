package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wediin/curator/lib/common"
)

const (
	formFieldFile        = "file"
	formFieldContributor = "contributor"
	photoStoreName       = "photos"
	defaultContributor   = "defaultContributor"
)

type UploadController struct {
	StorePath string
}

func (ctr *UploadController) PostController(c *gin.Context) {
	file, handler, err := c.Request.FormFile(formFieldFile)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	photoStorePath := ctr.StorePath + "/" + photoStoreName
	if err := common.CreateDir(photoStorePath); err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	contributor := c.Request.FormValue(formFieldContributor)
	if len(contributor) == 0 {
		contributor = defaultContributor
	}

	saveFilePath := genSaveFilePath(handler.Filename, photoStorePath, contributor)
	out, err := os.Create(saveFilePath)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "upload successful")
}

func genSaveFilePath(filename string, dir string, contributor string) string {
	saveFileName := fmt.Sprintf(
		"%s/%s-%d-%s",
		dir,
		contributor,
		time.Now().Unix(),
		filename,
	)
	return saveFileName
}
