package model

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Photo struct {
	ID          string       `json:"id"`
	Contributor string       `json:"contributor"`
	OriginURL   string       `json:"originURL"`
	ThumbURL    string       `json:"thumbURL"`
	WebviewURL  string       `json:"webviewURL"`
	Time        graphql.Time `json:"time"`
	Masked      bool         `json:"masked"`
}
