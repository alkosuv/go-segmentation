package sys

import (
	"os"
	"strings"
)

func Mkdir(path string, perm os.FileMode) error {
	if ok, err := IsExistDir(path); ok {
		return nil
	} else if err != nil {
		return err
	}

	if err := os.MkdirAll(path, perm); err != nil {
		return err
	}

	return nil
}

func IsExistDir(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Сtdir(path string) string {
	elems := strings.Split(path, "/")
	index := strings.LastIndex(elems[len(elems)-1], ".")
	if index < 0 {
		return path
	}

	return strings.Join(elems[:len(elems)-1], "/")
}

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	path := Сtdir(name)

	if err := Mkdir(path, perm); err != nil {
		return nil, err
	}

	return os.OpenFile(name, flag, perm)
}
