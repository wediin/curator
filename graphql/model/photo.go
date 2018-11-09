package model

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Photo struct {
	ID          string       `json:"id"`
	Contributor string       `json:"contributor"`
	Origin      *PhotoURL    `json:"origin"`
	Thumb       *PhotoURL    `json:"thumb"`
	Webview     *PhotoURL    `json:"webview"`
	Time        graphql.Time `json:"time"`
	Masked      bool         `json:"masked"`
}
