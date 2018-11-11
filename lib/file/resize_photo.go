package file

import (
	"image"
	"mime/multipart"

	"github.com/disintegration/imaging"
)

func ResizePhoto(imgFile multipart.File, targetPath string, maxLen int) error {
	photoWidth, photoHeight, err := GetPhotoSize(imgFile)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	if photoWidth < maxLen && photoHeight < maxLen {
		// do nothing
	} else if photoWidth > photoHeight {
		img = imaging.Resize(img, maxLen, 0, imaging.Lanczos)
	} else {
		img = imaging.Resize(img, 0, maxLen, imaging.Lanczos)
	}

	if err = imaging.Save(img, targetPath); err != nil {
		return err
	}

	if _, err := imgFile.Seek(0, 0); err != nil {
		return err
	}

	return nil
}
