package CIScriptsRuby

import (
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type RakeTest struct{}

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
