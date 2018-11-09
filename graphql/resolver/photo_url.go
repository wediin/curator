package resolver

import (
	"github.com/wediin/curator/graphql/model"
)

type photoURLResolver struct {
	photoURL *model.PhotoURL
}

func (r *photoURLResolver) Width() *int32 {
	return &r.photoURL.Width
}

func (r *photoURLResolver) Height() *int32 {
	return &r.photoURL.Height
}

func (r *photoURLResolver) URL() *string {
	return &r.photoURL.URL
}
