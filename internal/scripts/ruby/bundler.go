package CIScriptsRuby

import (
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type Bundler struct{}

func (b *Bundler) Run() error {
	if !c.CheckBinary("bundler") {
		c.Command("gem", "install", "bundler", "--no-ri", "--no-rdoc")
	}
	installPath, _ := c.ConfigFetch("ruby.bundler.install_path", "vendor")

	if !c.TestCommand("bundler", "check") {
		c.Command("bundler", "install", "--path", installPath)
	}
	return nil
}
