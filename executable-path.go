package helpers

import (
	"os"
	"path/filepath"
)

func GetWorkingDirectory() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return path, nil
}
