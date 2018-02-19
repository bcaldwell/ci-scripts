package CIScriptsRuby

import (
	"bytes"

	c "gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/CIScriptsHelpers"
)

type PublishGem struct{}

// out = capture_command("bundle", "exec", "rake", "-T")

// if out.include?("rake release[remote]")
// 		command("bundle", "exec", "rake", "release")
// else
// 		log_error("bundler/gem_tasks is required to use this script")
// end

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
