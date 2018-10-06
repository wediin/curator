package photo

import (
	"github.com/graphql-go/graphql"
	"github.com/wediin/curator/lib/google"
)

type Photo struct {
	ID              string
	Contributor     string
	Urls            []string
	CreateTimestamp string
	Masked          bool
}

func photosResolve(p graphql.ResolveParams) (interface{}, error) {
	mediaItems, err := google.RetrieveSharedPhotos()
	if err != nil {
		print("fail to retrieve share photos")
	}
	var sharePhotos []Photo
	for _, item := range mediaItems {
		sharePhotos = append(sharePhotos, Photo{ID: item.Id, Contributor: "elliottDefault", Urls: []string{item.BaseUrl}, CreateTimestamp: item.MediaMetadata.CreationTime, Masked: false})
	}
	return sharePhotos, nil
}
