package resolver

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/wediin/curator/graphql/model"
)

// Query
type photoArgs struct {
	ID string
}

func (r *Resolver) Photo(ctx context.Context, args photoArgs) (*photoResolver, error) {
	photo, err := r.PhotoClient.FindByID(args.ID)
	if err != nil {
		return nil, fmt.Errorf("photo: Fail to select photo by ID=[%v], err: (%v)", args.ID, err)
	}

	return &photoResolver{
		photo: &model.Photo{
			ID:          photo.ID.Hex(),
			Contributor: photo.Contributor,
			Origin: &model.PhotoURL{
				Width:  photo.Origin.Width,
				Height: photo.Origin.Height,
				URL:    photo.Origin.URL,
			},

			Thumb: &model.PhotoURL{
				Width:  photo.Thumb.Width,
				Height: photo.Thumb.Height,
				URL:    photo.Thumb.URL,
			},
			Webview: &model.PhotoURL{
				Width:  photo.Webview.Width,
				Height: photo.Webview.Height,
				URL:    photo.Webview.URL,
			},
			Time:   graphql.Time{photo.Time},
			Masked: photo.Masked,
		},
	}, nil
}

func (r *Resolver) Photos(ctx context.Context) (*[]*photoResolver, error) {
	photos, err := r.PhotoClient.Find()
	if err != nil {
		return nil, fmt.Errorf("photo: Fail to select photos, err: (%v)", err)
	}

	photoResolvers := make([]*photoResolver, 0)
	for i, _ := range photos {
		photo := photos[len(photos)-i-1]
		photoResolvers = append(photoResolvers, &photoResolver{
			photo: &model.Photo{
				ID:          photo.ID.Hex(),
				Contributor: photo.Contributor,
				Origin: &model.PhotoURL{
					Width:  photo.Origin.Width,
					Height: photo.Origin.Height,
					URL:    photo.Origin.URL,
				},

				Thumb: &model.PhotoURL{
					Width:  photo.Thumb.Width,
					Height: photo.Thumb.Height,
					URL:    photo.Thumb.URL,
				},
				Webview: &model.PhotoURL{
					Width:  photo.Webview.Width,
					Height: photo.Webview.Height,
					URL:    photo.Webview.URL,
				},
				Time:   graphql.Time{photo.Time},
				Masked: photo.Masked,
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

func (r *photoResolver) Origin() *photoURLResolver {
	return &photoURLResolver{
		photoURL: r.photo.Origin,
	}
}

func (r *photoResolver) Thumb() *photoURLResolver {
	return &photoURLResolver{
		photoURL: r.photo.Thumb,
	}
}

func (r *photoResolver) Webview() *photoURLResolver {
	return &photoURLResolver{
		photoURL: r.photo.Webview,
	}
}

func (r *photoResolver) Time() *graphql.Time {
	return &r.photo.Time
}

func (r *photoResolver) Masked() *bool {
	return &r.photo.Masked
}
