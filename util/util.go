package util

import (
	"os"
	"path"
	"strings"
)

var extension = ".dapi.json"

// ProjectPath returns the path to the project file for the supplied name
func ProjectPath(name string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(name, extension) {
		name += extension
	}
	return path.Join(cwd, name), nil
}

// Must panics if err is not nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
