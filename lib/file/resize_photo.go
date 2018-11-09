package file

import (
	"image"
	"mime/multipart"

	"github.com/disintegration/imaging"
)

func ResizePhoto(file multipart.File, targetPath string, width int, height int) error {
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	img = imaging.Resize(img, width, height, imaging.Lanczos)
	if err = imaging.Save(img, targetPath); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	return nil
}
