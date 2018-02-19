package CIScriptsRuby

import (
	c "gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/CIScriptsHelpers"
)

type RakeTest struct{}

// if test_command?("bundler", "exec", "rake", "-V")
// return command("bundler", "exec", "rake", "test")
// end

// unless installed?("rake")
// command("gem", "install", "rake")
// end
// command("rake", "test")

func (r *RakeTest) Run() error {
	// Use bundler version if installed
	if !c.TestCommand("bundler", "exec", "rake", "-V") {
		c.Command("bundler", "exec", "rake", "test")
		return nil
	}

	if !c.CheckBinary("rake") {
		c.Command("gem", "install", "rake")
	}

	c.Command("rake", "test")

	return nil
}
