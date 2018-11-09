package db

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type PhotoModel struct {
	ID          objectid.ObjectID `bson:"_id"`
	Contributor string            `bson: "contributor"`
	Origin      PhotoURLModel     `bson: "origin"`
	Thumb       PhotoURLModel     `bson: "thumb"`
	Webview     PhotoURLModel     `bson: "webview"`
	Time        time.Time         `bson: "time"`
	Masked      bool              `bson: "masked"`
}

type PhotoURLModel struct {
	Width  int32  `bson:"width"`
	Height int32  `bson:"height"`
	URL    string `bson:"url"`
}
