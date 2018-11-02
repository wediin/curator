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
	ThumbWidth     int
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

	thumbF, _, err := c.Request.FormFile(formFieldFile)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}
	defer thumbF.Close()

	if err = file.ResizePhoto(thumbF, thumbFilePath, ctr.ThumbWidth, 0); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	webViewF, _, err := c.Request.FormFile(formFieldFile)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}
	defer webViewF.Close()

	photoWidth, photoHeight, err := file.GetPhotoSize(webViewF)
	if err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	if _, err := webViewF.Seek(0, 0); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	if ctr.WebviewMaxLen > photoWidth && ctr.WebviewMaxLen > photoHeight {
		if err = file.SaveFile(webViewF, webViewFilePath); err != nil {
			statusError(c, http.StatusInternalServerError, err)
			return
		}
	} else if photoWidth > photoHeight {
		if err = file.ResizePhoto(webViewF, webViewFilePath, ctr.WebviewMaxLen, 0); err != nil {
			statusError(c, http.StatusInternalServerError, err)
			return
		}
	} else {
		if err = file.ResizePhoto(webViewF, webViewFilePath, 0, ctr.WebviewMaxLen); err != nil {
			statusError(c, http.StatusInternalServerError, err)
			return
		}
	}

	photo := db.PhotoModel{
		ID:          id,
		Contributor: contributor,
		OriginURL:   common.JoinURL(ctr.Url, ctr.PhotoRouter, originPath, photoFileName),
		ThumbURL:    common.JoinURL(ctr.Url, ctr.PhotoRouter, thumbPath, photoFileName),
		WebviewURL:  common.JoinURL(ctr.Url, ctr.PhotoRouter, webViewPath, photoFileName),
		Time:        time.Now(),
		Masked:      false,
	}

	if err = ctr.PhotoClient.Insert(&photo); err != nil {
		statusError(c, http.StatusInternalServerError, err)
		return
	}

	statusOK(c, http.StatusOK, "upload successful")
}
