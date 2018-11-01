package file

import (
	"image"
	"io"

	"github.com/disintegration/imaging"
)

func SaveAsThumb(src io.Reader, thumbFilePath string, width int) error {
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	img = imaging.Resize(img, width, 0, imaging.Lanczos)
	if err = imaging.Save(img, thumbFilePath); err != nil {
		return err
	}
	return nil
}
