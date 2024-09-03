package file

import (
	"os"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Create(filename string) (*os.File, error) {
	return os.Create(filename)
}
