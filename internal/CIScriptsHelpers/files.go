package CIScriptsHelpers

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

func CopyAndExpandFolder(srcFolder, destFolder string) error {
	err := filepath.Walk(srcFolder, func(src string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		srcFile, err := os.Stat(src)

		// If no error
		if err != nil {
			return err
		}

		// File & Folder Mode
		srcfMode := srcFile.Mode()

		if srcfMode.IsRegular() {
			relPath, err := filepath.Rel(srcFolder, src)
			if err != nil {
				return err
			}

			dest := filepath.Join(destFolder, relPath)
			// Append to Files Array
			f, err := os.Create(dest)
			if err != nil {
				return err
			}

			if err = os.Chmod(f.Name(), srcfMode.Perm()); err != nil {
				return err
			}

			s, err := ioutil.ReadFile(src)
			if err != nil {
				return err
			}

			// expand env but don't change the value if the env variable doesn't exist
			expandedSrc := os.Expand(string(s), func(s string) string {
				var expandedVal string
				var ok bool

				if expandedVal, ok = os.LookupEnv(s); !ok {
					return fmt.Sprintf("${%s}", s)
				}
				return expandedVal
			})

			_, err = io.Copy(f, strings.NewReader(expandedSrc))
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
