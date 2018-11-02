package file

import (
	"image"
	"io"

	"github.com/disintegration/imaging"
)

func ResizePhoto(src io.Reader, targetPath string, width int, height int) error {
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	img = imaging.Resize(img, width, height, imaging.Lanczos)
	if err = imaging.Save(img, targetPath); err != nil {
		return err
	}
	return nil
}
