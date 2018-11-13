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
	defaultContributor   = "default"
	originPath           = "origin"
	thumbPath            = "thumb"
	webViewPath          = "webview"
)

type UploadController struct {
	Url            string
	PhotoStorePath string
	PhotoRouter    string
	PhotoDir       string
	PhotoClient    *db.PhotoClient
	ThumbMaxLen    int
	WebviewMaxLen  int
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
	webViewStorePath := filepath.Join(ctr.PhotoStorePath, webViewPath)
	if err := common.CreateMultiDir(originStorePath, thumbStorePath, webViewStorePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	contributor := c.Request.FormValue(formFieldContributor)
	if len(contributor) == 0 {
		contributor = defaultContributor
	}

	id := objectid.New()

	photoFileName := fmt.Sprintf("%s-%s-%s", contributor, id.Hex(), handler.Filename)
	originFilePath := filepath.Join(originStorePath, photoFileName)
	thumbFilePath := filepath.Join(thumbStorePath, photoFileName)
	webViewFilePath := filepath.Join(webViewStorePath, photoFileName)

	if err = file.SaveFile(f, originFilePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	if err = file.SaveFile(f, thumbFilePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	if err = file.SaveFile(f, webViewFilePath); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	photoWidth, photoHeight, err := file.GetPhotoSize(f)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	photo := db.PhotoModel{
		ID:          id,
		Contributor: contributor,
		Origin: db.PhotoURLModel{
			Width:  int32(photoWidth),
			Height: int32(photoHeight),
			URL:    common.JoinURL(ctr.Url, ctr.PhotoRouter, originPath, photoFileName),
		},
		Thumb: db.PhotoURLModel{
			Width:  int32(photoWidth),
			Height: int32(photoHeight),
			URL:    common.JoinURL(ctr.Url, ctr.PhotoRouter, thumbPath, photoFileName),
		},
		Webview: db.PhotoURLModel{
			Width:  int32(photoWidth),
			Height: int32(photoHeight),
			URL:    common.JoinURL(ctr.Url, ctr.PhotoRouter, webViewPath, photoFileName),
		},
		Time:   time.Now(),
		Masked: false,
	}

	if err = ctr.PhotoClient.Insert(&photo); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	statusOK(c, http.StatusOK, "upload successful")
}
