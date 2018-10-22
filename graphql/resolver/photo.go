package resolver

import (
	"context"
	"fmt"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/wediin/curator/graphql/model"
	"github.com/wediin/curator/lib/db"
)

// Query
func (r *Resolver) Photos(ctx context.Context) (*[]*photoResolver, error) {
	client, err := db.NewClient(r.MongoServer)
	if err != nil {
		return nil, fmt.Errorf("photo: Fail to new db client, err: (%v)", err)
	}

	photos, err := client.SelectPhotos(r.MongoDB, r.PhotoMongoCollection)
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
				Timestamp:   strconv.FormatInt(photo.Time.Unix(), 10),
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

func (r *photoResolver) Timestamp() *string {
	return &r.photo.Timestamp
}

func (r *photoResolver) Masked() *bool {
	return &r.photo.Masked
}
