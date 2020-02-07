package CIScripts

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
	CIScriptsDocker "github.com/bcaldwell/ci-scripts/internal/scripts/docker"
	CIScriptsGit "github.com/bcaldwell/ci-scripts/internal/scripts/git"
	CIScriptsGithub "github.com/bcaldwell/ci-scripts/internal/scripts/github"
	CIScriptsGo "github.com/bcaldwell/ci-scripts/internal/scripts/go"
	CIScriptsKubernetes "github.com/bcaldwell/ci-scripts/internal/scripts/kubernetes"
	CIScriptsRuby "github.com/bcaldwell/ci-scripts/internal/scripts/ruby"
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

	"git/files_changed": &CIScriptsGit.FilesChanged{},

	"docker/deploy":               &CIScriptsDocker.BuildAndDeploy{},
	"docker/build_and_push_image": &CIScriptsDocker.BuildAndPushImage{},

	"github/release":           &CIScriptsGithub.Release{},
	"github/release_checksums": &CIScriptsGithub.ReleaseChecksums{},

	"kubernetes/deploy": &CIScriptsKubernetes.Deploy{},
}

func Execute() {
	if len(os.Args) <= 1 {
		CIScriptsHelpers.LogError("No scripts specified")
		os.Exit(0)
	}

	CIScriptsHelpers.ParseCLIArguments()

	scriptName := os.Args[1]
	if curScript, ok := scripts[scriptName]; ok {
		err := curScript.Run()
		if err != nil {
			CIScriptsHelpers.LogError("Error in script %s: %s", scriptName, err)
			os.Exit(1)
		}
	} else if scriptName == "list" {
		keys := make([]string, 0, len(scripts))
		for k := range scripts {
			keys = append(keys, k)
		}
		fmt.Println(strings.Join(keys, "\n"))
	} else {
		CIScriptsHelpers.LogError("Script %s not found\n", scriptName)
	}
}
