package CIScriptsRuby

import (
	c "gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/CIScriptsHelpers"
)

type Bundler struct{}

// unless installed?("bundler")
// command("gem", "install", "bundler", "--no-ri", "--no-rdoc")
// end

// install_path = env_fetch("BUNDLER_INSTALL_PATH", "vendor")

// unless test_command?("bundler", "check")
// command("bundler", "install", "--path", install_path)
// end

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
