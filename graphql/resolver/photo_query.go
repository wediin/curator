package resolver

import (
	"context"

	"github.com/wediin/curator/graphql/model"
	"github.com/wediin/curator/lib/google"
)

func (r *Resolver) Photos(ctx context.Context) (*[]*photoResolver, error) {
	mediaItems, err := google.RetrieveSharedPhotos()
	if err != nil {
		return nil, err
	}

	photos := make([]*photoResolver, 0)
	for _, item := range mediaItems {
		photos = append(photos, &photoResolver{
			photo: &model.Photo{
				ID: item.Id,
				Contributor: "elliottDefault",
				Urls: []string{item.BaseUrl},
				Timestamp: item.MediaMetadata.CreationTime,
				Masked: false,
			},
		})
	}

	return &photos, nil
}
