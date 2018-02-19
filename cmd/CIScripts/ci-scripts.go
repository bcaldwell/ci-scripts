package CIScripts

import (
	"os"

	"gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/CIScriptsHelpers"
	"gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/scripts/github"
	"gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/scripts/go"
	"gitlab.caldwellbenjam.in/benjamin/ci-scripts/internal/scripts/ruby"
)

type script interface {
	Run() error
}

var scripts = map[string]script{
	"ruby/bundler":     &CIScriptsRuby.Bundler{},
	"ruby/rake_test":   &CIScriptsRuby.RakeTest{},
	"ruby/rubocop":     &CIScriptsRuby.Rubocop{},
	"ruby/publish_gem": &CIScriptsRuby.PublishGem{},

	"go/build": &CIScriptsGo.Build{},

	"github/release": &CIScriptsGithub.Release{},
}

func Execute() {
	if len(os.Args) <= 1 {
		CIScriptsHelpers.LogError("No scripts specified")
	}
	for _, scriptName := range os.Args[1:] {
		if curScript, ok := scripts[scriptName]; ok {
			curScript.Run()
		} else {
			CIScriptsHelpers.LogError("Script %s not found\n", scriptName)
		}
	}
}
