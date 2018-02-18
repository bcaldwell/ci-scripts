package main

import (
	"gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/scripts/ruby"
)

func main() {

	b := CIScriptsRuby.Bundler{}
	b.Run()
}
