package CIScriptsGo

import (
	"fmt"
	"strings"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type Build struct {
	buildCommand []string
}

func (b *Build) Run() error {

	if !c.CheckBinary("gox") {
		c.Command("go", "get", "-u", "github.com/mitchellh/gox")
	}

	b.buildCommand = []string{"gox"}

	b.addBuildOption("ldflags")
	b.addBuildOption("parallel")
	b.addBuildOption("osarch")
	b.addBuildOption("os")
	b.addBuildOption("output")

	if args, ok := c.ConfigFetch("go.build.go_build_args"); ok {
		b.buildCommand = append(b.buildCommand, strings.Split(args, " ")...)
	}

	err := c.Command(b.buildCommand...)
	if err != nil {
		return err
	}
	return nil
}

func (b *Build) addBuildOption(option string) {
	configName := "go.build." + option
	if val, ok := c.ConfigFetch(configName); ok {
		b.buildCommand = append(b.buildCommand, fmt.Sprintf("-%s=%s", option, val))
	}
}
