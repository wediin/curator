package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/wediin/curator/graphql/model"
)

type photoResolver struct {
	photo *model.Photo
}

func (r *photoResolver) ID() *graphql.ID {
	id := graphql.ID(r.photo.ID)
	return &id
}

func (r *photoResolver) Contributor() *string {
	return &r.photo.Contributor
}

func (r *photoResolver) Urls() *[]*string {
	urls := make([]*string, len(r.photo.Urls))

	for i := range urls {
		urls[i] = &r.photo.Urls[i]
	}

	return &urls
}

func (r *photoResolver) Timestamp() *string {
	return &r.photo.Timestamp
}

func (r *photoResolver) Masked() *bool {
	return &r.photo.Masked
}

