package CIScriptsHelpers

import (
	"os"
	"path/filepath"
	"strings"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func ExpandPath(path string) string {
	home := os.Getenv("HOME")

	if path == "~" {
		// In case of "~", which won't be caught by the "else if"
		path = home
	} else if strings.HasPrefix(path, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		path = filepath.Join(home, path[2:])
	}

	return os.ExpandEnv(path)
}

func RecursiveFilesInFolder(folder string) ([]string, error) {
	files := make([]string, 0)

	err := filepath.Walk(folder, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		f, err := os.Stat(path)

		// If no error
		if err != nil {
			return err
		}

		// File & Folder Mode
		fMode := f.Mode()

		if fMode.IsRegular() {
			// Append to Files Array
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
