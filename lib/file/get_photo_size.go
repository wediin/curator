package file

import (
	"image"
	"io"
)

func GetPhotoSize(src io.Reader) (int, int, error) {
	config, _, err := image.DecodeConfig(src)
	if err != nil {
		return -1, -1, err
	}

	return config.Width, config.Height, nil
}
