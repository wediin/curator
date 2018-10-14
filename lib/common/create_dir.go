package common

import (
	"os"
)

func CreateDir(dir string) error {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		return nil
	}

	return err
}
