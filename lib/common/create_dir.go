package common

import (
	"os"
)

func CreateMultiDir(dirs ...string) error {
	for _, dir := range dirs {
		if err := CreateDir(dir); err != nil {
			return err
		}
	}
	return nil
}

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
