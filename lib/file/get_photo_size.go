package file

import (
	"mime/multipart"
	"os"

	"github.com/disintegration/imageorient"
)

func GetPhotoSize(file multipart.File) (int, int, error) {
	config, _, err := imageorient.DecodeConfig(file)
	if err != nil {
		return -1, -1, err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return -1, -1, err
	}

	return config.Width, config.Height, nil
}

func GetPhotoSizeByPath(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, -1, err
	}
	return GetPhotoSize(file)
}
