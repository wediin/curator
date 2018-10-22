package model

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type Photo struct {
	ID          string       `json:"id"`
	Contributor string       `json:"contributor"`
	Urls        []string     `json:"urls"`
	Time        graphql.Time `json:"time"`
	Masked      bool         `json:"masked"`
}
