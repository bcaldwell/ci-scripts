package CIScriptsRuby

import (
	c "gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/CIScriptsHelpers"
)

type Rubocop struct{}

// unless installed?("rubocop")
// command("gem", "install", "rubocop")
// end
// # ugh check bundle rubocop as well
// # puts test_command?("bundle exec rubocop")

// command("rubocop")

func (b *Rubocop) Run() error {
	if !c.TestCommand("bundler", "exec", "rubocop", "-V") {
		c.Command("bundler", "exec", "rubocop")
		return nil
	}

	if !c.CheckBinary("rake") {
		c.Command("gem", "install", "rubocop")
	}

	c.Command("rubocop")
	return nil
}
