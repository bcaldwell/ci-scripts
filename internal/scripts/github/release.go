package CIScriptsGithub

import (
	"fmt"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type Release struct {
	command []string
}

// ghr -u bcaldwell -r devctl $BUILD_VERSION dist/release/

func (b *Release) Run() error {

	c.ConfigSetDefault("github.release.path", ".")
	c.RequiredConfigKey("github.release.tag")

	checksumsScripts := ReleaseChecksums{}

	err := checksumsScripts.Run()
	if err != nil {
		c.LogError("Error in script github/release_checksums: %s", err)
		// return nil cause error already logged here
		return nil
	}

	if !c.CheckBinary("ghr") {
		c.Command("go", "get", "-u", "github.com/tcnksm/ghr")
	}

	b.command = []string{"ghr"}

	b.addBuildOption("t", "token")
	b.addBuildOption("u", "username")
	b.addBuildOption("r", "repo")
	b.addBuildOption("c", "commit")
	b.addBuildOption("b", "body")
	b.addBuildOption("p", "parallel")

	if _, ok := c.ConfigFetch("github.release.delete"); ok {
		b.command = append(b.command, "-delete")
	}

	if _, ok := c.ConfigFetch("github.release.draft"); ok {
		b.command = append(b.command, "-draft")
	}

	if tag, ok := c.ConfigFetch("github.release.tag"); ok {
		b.command = append(b.command, tag)
	}

	if path, ok := c.ConfigFetch("github.release.path"); ok {
		b.command = append(b.command, path)
	}

	return c.Command(b.command...)
}

func (b *Release) addBuildOption(flag, option string) {
	configName := "github.release." + option
	if val, ok := c.ConfigFetch(configName); ok {
		b.command = append(b.command, fmt.Sprintf("-%s", flag), val)
	}
}
