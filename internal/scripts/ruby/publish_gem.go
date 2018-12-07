package CIScriptsRuby

import (
	"bytes"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type PublishGem struct{}

func (b *PublishGem) Run() error {
	out, err := c.CaptureCommand("bundle", "exec", "rake", "-T")
	if err != nil {
		return err
	}

	if bytes.Contains(out, []byte("rake release[remote]")) {
		c.Command("bundle", "exec", "rake", "release")
	} else {
		c.LogError("bundler/gem_tasks is required to use this script")
	}

	return nil
}
