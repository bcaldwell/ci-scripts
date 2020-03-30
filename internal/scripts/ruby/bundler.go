package CIScriptsRuby

import (
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type Bundler struct{}

func (b *Bundler) Run() error {
	if !c.CheckBinary("bundler") {
		err := c.Command("gem", "install", "bundler", "--no-ri", "--no-rdoc")
		if err != nil {
			return err
		}
	}

	installPath, _ := c.ConfigFetch("ruby.bundler.install_path", "vendor")

	if !c.TestCommand("bundler", "check") {
		err = c.Command("bundler", "install", "--path", installPath)
		if err != nil {
			return err
		}
	}

	return nil
}
