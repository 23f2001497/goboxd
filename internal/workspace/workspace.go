package workspace

import "os"

func Create() (string, error) {
	dir, err := os.MkdirTemp("", "goboxd-*")
	if err != nil {
		return "", err
	}
	return dir, nil
}

func Cleanup(dir string) {
	os.RemoveAll(dir)
}
