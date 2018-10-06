package google

import (
	"log"

	"google.golang.org/api/photoslibrary/v1"
)

const (
	RequestPageSize = 50
	// FIXME: decide a place to store this id
	ShareAlbumId = "ADOn-aBQdy0p01WhG8qAU5nKESq6g-Ln0R9IC4KOuBV1atJJFyUnUYDVBjEMcfg0Oyfc70DwDtgK"
)

func RetrieveSharedPhotos() ([]*photoslibrary.MediaItem, error) {
	//FIXME: Try different scope for this function
	srv, err := GetPhotosService(PhotosFullScope)
	if err != nil {
		log.Fatalf("Fail to get photo service.")
		return nil, err
	}

	shareAlbum := photoslibrary.Album{Id: ShareAlbumId}
	mediaItems, err := listSpecificAlbum(srv, &shareAlbum)
	if err != nil {
		log.Fatalf("Fail to get list of share album.")
		return nil, err
	}
	return mediaItems, nil
}

// List specific album
func listSpecificAlbum(srv *photoslibrary.Service, a *photoslibrary.Album) ([]*photoslibrary.MediaItem, error) {
	req := &photoslibrary.SearchMediaItemsRequest{PageSize: RequestPageSize, AlbumId: a.Id}
	mediaItems, err := executeSearchMediaItemsRequest(srv, req)
	return mediaItems, err
}

func executeSearchMediaItemsRequest(srv *photoslibrary.Service, req *photoslibrary.SearchMediaItemsRequest) ([]*photoslibrary.MediaItem, error) {
	hasMore := true
	var mediaItems []*photoslibrary.MediaItem
	for hasMore {
		items, err := srv.MediaItems.Search(req).Do()
		if err != nil {
			return nil, err
		}
		mediaItems = append(mediaItems, items.MediaItems...)
		req.PageToken = items.NextPageToken
		if req.PageToken == "" {
			hasMore = false
		}
	}
	return mediaItems, nil
}
