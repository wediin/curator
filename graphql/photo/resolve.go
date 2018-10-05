package photo

import (
	"github.com/graphql-go/graphql"
)

func photosResolve(p graphql.ResolveParams) (interface{}, error) {
	return photos, nil
}
