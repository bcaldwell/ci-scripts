package CIScriptsRuby

import (
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type Rubocop struct{}

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
