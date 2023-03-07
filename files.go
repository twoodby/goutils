package goutils

import (
	"errors"
	"os"
)

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func FileExistStrict(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
