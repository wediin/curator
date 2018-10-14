package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	UploadFolder       = "uploadFolder"
	FormFileField      = "file"
	ContributorField   = "contributor"
	DefaultContributor = "defaultContributor"
)

type UploadController struct{}

func (ctr *UploadController) PostController(c *gin.Context) {
	file, handler, err := c.Request.FormFile(FormFileField)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	saveFileName := genSaveFileName(handler.Filename, c.Request)
	out, err := os.Create(saveFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	io.Copy(out, file)

	c.String(http.StatusOK, "upload successful\n")
}

func genSaveFileName(filename string, r *http.Request) string {
	contributor := r.FormValue(ContributorField)
	if len(contributor) == 0 {
		contributor = DefaultContributor
	}
	saveFileName := fmt.Sprintf(
		"./%s/%s-%d-%s",
		UploadFolder,
		contributor,
		time.Now().Unix(),
		filename,
	)
	return saveFileName

}
