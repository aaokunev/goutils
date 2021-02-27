package file

import (
	"github.com/pkg/errors"
	"io"
	"os"
)

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func IsFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func IsDirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func MustSeekToStart(seeker io.Seeker) {
	_, err := seeker.Seek(0, 0)

	if err != nil {
		panic(errors.Wrap(err, "unable to seek to start"))
	}
}
