package photo

import (
	"github.com/graphql-go/graphql"
)

func GetPhotosQueryField() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(photoType),
		Resolve: photosResolve,
	}
}
