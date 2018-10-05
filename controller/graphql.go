package controller

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/wediin/curator/graphql/photo"
)

type GraphqlController struct{}

func (ctr *GraphqlController) newSchema() (graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"photos": photo.GetPhotosQueryField(),
			},
		}),
	})

	return schema, err
}

func (ctr *GraphqlController) NewHandler() (*handler.Handler) {
	schema, err := ctr.newSchema()

	if err != nil {
		log.Fatalf("Fail to create schema, error: %v", err)
		return nil
	}

	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
}
