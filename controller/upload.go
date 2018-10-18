package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wediin/curator/lib/common"
	"github.com/wediin/curator/lib/db"
	"github.com/wediin/curator/lib/file"
	"gopkg.in/mgo.v2/bson"
)

const (
	formFieldFile        = "file"
	formFieldContributor = "contributor"
	defaultContributor   = "defaultContributor"
)

type UploadController struct {
	Url                  string
	MongoServer          string
	MongoDB              string
	PhotoMongoCollection string
	PhotoStorePath       string
	PhotoRouter          string
	PhotoDir             string
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

	id := bson.NewObjectId()

	photoFileName := fmt.Sprintf("%s-%s-%s", contributor, id.Hex(), handler.Filename)
	photoFilePath := fmt.Sprintf("%s/%s", ctr.PhotoStorePath, photoFileName)

	err = file.SaveFile(f, photoFilePath)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	client, err := db.NewClient(ctr.MongoServer)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	photo := db.ModelPhoto{
		Contributor: contributor,
		Urls: []string{
			ctr.Url + ctr.PhotoRouter + "/" + photoFileName,
		},
		Time: time.Now(),
		Mask: false,
	}
	photo.AssignId(id)

	err = client.Insert(ctr.MongoDB, ctr.PhotoMongoCollection, photo)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	statusOK(c, http.StatusOK, "upload successful")
}
