package photo

import (
	"github.com/graphql-go/graphql"
)

type photo struct {
	ID            string   `json:"id"`
	Contributor   string   `json:"contributor"`
	Urls          []string `json:"urls"`
}

// FIXME: Mockup data
var photos []photo = []photo{
	photo{
		ID: "5b9d0aeb92cf31ce0d8040df",
		Contributor: "Kevin",
		Urls: []string{
			"https://drive.google.com/some-hash-tag",
		},
	},
}

var photoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Photo",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.String},
		"contributor":  &graphql.Field{Type: graphql.String},
		"urls":         &graphql.Field{Type: graphql.NewList(graphql.String)},
	},
})
