package CIScriptsGo

import (
	"fmt"
	"strings"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type Build struct {
	buildCommand []string
}

// def run
// command("go", "get", "github.com/mitchellh/gox")

// build_command = ["gox"]

// build_command = add_option("ldflags", build_command)

// parallel = env_fetch("GO_BUILD_PARALLEL", "2")
// build_command.push("-parallel")
// build_command.push(parallel)

// build_command = add_option("osarch", build_command)
// build_command = add_option("output", build_command)

// if gox_args = env_fetch("GO_BUILD_ARGS")
// 		build_command.concat(gox_args.split)
// end

// # idk why this is needed...
// command(build_command.join(" "))
// end

// private

// def add_option(name, build_command)
// env_name = "GO_BUILD_#{name.upcase}"
// if val = env_fetch(env_name)
// 		build_command.push("-#{name}")
// 		build_command.push("\"#{val}\"")
// end
// build_command
// end

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
