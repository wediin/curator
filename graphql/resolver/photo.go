package resolver

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/wediin/curator/graphql/model"
	"github.com/wediin/curator/lib/common"
	"github.com/wediin/curator/lib/google"
)

// Query
func (r *Resolver) Photos(ctx context.Context) (*[]*photoResolver, error) {
	mediaItems, err := google.RetrieveSharedPhotos()
	if err != nil {
		return nil, err
	}

	photos := make([]*photoResolver, 0)
	for _, item := range mediaItems {
		photos = append(photos, &photoResolver{
			photo: &model.Photo{
				ID:          item.Id,
				Contributor: "DefaultContributor",
				Urls:        []string{item.BaseUrl},
				Timestamp:   common.TransformRFC3339Time(item.MediaMetadata.CreationTime),
				Masked:      false,
			},
		})
	}

	// append local files
	files, err := ioutil.ReadDir("uploadFolder")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		photos = append(photos, &photoResolver{
			photo: &model.Photo{
				ID:          strconv.FormatInt(time.Now().Unix(), 10),
				Contributor: "localfolder",
				Urls:        []string{"localhost:9527/usercontent/" + file.Name()},
				Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
				Masked:      false,
			},
		})
	}
	return &photos, nil
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
