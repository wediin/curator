package db

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type PhotoModel struct {
	ID          objectid.ObjectID `bson:"_id"`
	Contributor string            `bson: "contributor"`
	Urls        []string          `bson: "urls"`
	Time        time.Time         `bson: "time"`
	Masked      bool              `bson: "masked"`
}
