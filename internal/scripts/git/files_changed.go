package CIScriptsGit

import (
	"errors"
	"strings"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type FilesChanged struct {
}

func (b *FilesChanged) Run() error {
	changePatterns := strings.Split(c.RequiredConfigFetch("git.files_changed.prefix"), ",")
	if len(changePatterns) == 0 {
		c.LogError("No file patterns to check passed in as CLI arguments")
		return nil
	}

	result, err := c.CaptureCommand("git", "diff", "HEAD^1", "HEAD", "--name-only")
	if err != nil {
		return err
	}

	filesChanged := strings.Split(string(result), "\n")

	for _, pattern := range changePatterns {
		pattern = strings.TrimSpace(pattern)

		for _, file := range filesChanged {
			matched := strings.HasPrefix(file, pattern)
			if matched {
				return nil
			}
		}
	}

	return errors.New("no match")
}
