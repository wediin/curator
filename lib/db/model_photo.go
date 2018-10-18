package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ModelPhoto struct {
	_id         bson.ObjectId
	Contributor string
	Urls        []string
	Time        time.Time
	Mask        bool
}

func (m *ModelPhoto) AssignId(id bson.ObjectId) {
	m._id = id
}
