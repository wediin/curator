package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
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
	originPath           = "origin"
	thumbPath            = "thumb"
)

type UploadController struct {
	Url            string
	PhotoStorePath string
	PhotoRouter    string
	PhotoDir       string
	PhotoClient    *db.PhotoClient
	ThumbWidth     int
}

func (ctr *UploadController) PostController(c *gin.Context) {
	f, handler, err := c.Request.FormFile(formFieldFile)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}
	defer f.Close()

	originStorePath := filepath.Join(ctr.PhotoStorePath, originPath)
	thumbStorePath := filepath.Join(ctr.PhotoStorePath, thumbPath)
	if err := common.CreateMultiDir(originStorePath, thumbStorePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	contributor := c.Request.FormValue(formFieldContributor)
	if len(contributor) == 0 {
		contributor = defaultContributor
	}

	id := objectid.New()

	photoFileName := fmt.Sprintf("%s-%s-%s", contributor, id.Hex(), handler.Filename)
	photoFilePath := filepath.Join(originStorePath, photoFileName)
	thumbFilePath := filepath.Join(thumbStorePath, photoFileName)

	if err = file.SaveFile(f, photoFilePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	thumbF, _, err := c.Request.FormFile(formFieldFile)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}
	defer thumbF.Close()

	if err = file.SaveAsThumb(thumbF, thumbFilePath, ctr.ThumbWidth); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	photo := db.PhotoModel{
		ID:          id,
		Contributor: contributor,
		OriginURL:   common.JoinURL(ctr.Url, ctr.PhotoRouter, originPath, photoFileName),
		ThumbURL:    common.JoinURL(ctr.Url, ctr.PhotoRouter, thumbPath, photoFileName),
		Time:        time.Now(),
		Masked:      false,
	}

	if err = ctr.PhotoClient.Insert(&photo); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	statusOK(c, http.StatusOK, "upload successful")
}
