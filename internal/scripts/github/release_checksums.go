package CIScriptsGithub

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type ReleaseChecksums struct{}

func (r *ReleaseChecksums) Run() error {

	c.TimedRun("Generating SHAs", func() error {
		c.ConfigSetDefault("github.release.path", ".")
		outputFolder, _ := c.ConfigFetch("github.release.path")

		shaFile, _ := c.ConfigFetch("github.release.checksum_file", "checksums.txt")

		shaPath := path.Join(outputFolder, shaFile)
		fmt.Println(shaPath)
		output, err := os.Create(shaPath)
		if err != nil {
			return err
		}
		defer output.Close()

		err = filepath.Walk(outputFolder, func(releasePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			filename := path.Base(releasePath)

			if info.IsDir() || filename == shaFile || strings.HasPrefix(filename, ".") {
				return nil
			}

			var f io.Reader
			f, err = os.Open(releasePath)
			if err != nil {
				return err
			}

			h := sha256.New()
			_, err = io.Copy(h, f)
			if err != nil {
				return err
			}

			output.WriteString(fmt.Sprintf("%x\t%s\n", h.Sum(nil), filename))

			return nil
		})
		return nil
	})
	return nil
}
