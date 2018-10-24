package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/wediin/curator/graphql/model"
	"github.com/wediin/curator/lib/db"
)

// Query
func (r *Resolver) Photos(ctx context.Context) (*[]*photoResolver, error) {
	client, err := db.NewPhotoClient(r.MongoServer, r.MongoDB, r.PhotoMongoCollection)
	if err != nil {
		return nil, fmt.Errorf("photo: Fail to new photo db client, err: (%v)", err)
	}

	photos, err := client.Select()
	if err != nil {
		return nil, fmt.Errorf("photo: Fail to select photos, err: (%v)", err)
	}

	photoResolvers := make([]*photoResolver, 0)
	for _, photo := range photos {
		photoResolvers = append(photoResolvers, &photoResolver{
			photo: &model.Photo{
				ID:          photo.ID.Hex(),
				Contributor: photo.Contributor,
				Urls:        photo.Urls,
				Time:        graphql.Time{photo.Time},
				Masked:      photo.Masked,
			},
		})
	}
	return &photoResolvers, nil
}

// Resolver
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

func (r *photoResolver) Time() *graphql.Time {
	return &r.photo.Time
}

func (r *photoResolver) Masked() *bool {
	return &r.photo.Masked
}
