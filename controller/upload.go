package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/wediin/curator/lib/common"
	"github.com/wediin/curator/lib/db"
	"github.com/wediin/curator/lib/file"
)

const (
	formFieldFile        = "file"
	formFieldContributor = "contributor"
	defaultContributor   = "defaultContributor"
)

type UploadController struct {
	Url            string
	PhotoStorePath string
	PhotoRouter    string
	PhotoDir       string
	PhotoClient    *db.PhotoClient
}

func (ctr *UploadController) PostController(c *gin.Context) {
	f, handler, err := c.Request.FormFile(formFieldFile)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}
	defer f.Close()

	if err := common.CreateDir(ctr.PhotoStorePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	contributor := c.Request.FormValue(formFieldContributor)
	if len(contributor) == 0 {
		contributor = defaultContributor
	}

	id := objectid.New()

	photoFileName := fmt.Sprintf("%s-%s-%s", contributor, id.Hex(), handler.Filename)
	photoFilePath := fmt.Sprintf("%s/%s", ctr.PhotoStorePath, photoFileName)

	err = file.SaveFile(f, photoFilePath)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	photo := db.PhotoModel{
		ID:          id,
		Contributor: contributor,
		Urls: []string{
			ctr.Url + ctr.PhotoRouter + "/" + photoFileName,
		},
		Time:   time.Now(),
		Masked: false,
	}

	err = ctr.PhotoClient.Insert(&photo)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	statusOK(c, http.StatusOK, "upload successful")
}
