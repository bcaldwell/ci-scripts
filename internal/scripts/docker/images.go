package CIScriptsDocker

import (
	"flag"
	"fmt"
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
	"github.com/bcaldwell/sshtun"
	"os"
	"path"
	"strings"
)

func init() {
	flag.Var(&dockerBuildArg, "build-arg", "Docker build-arg.")
}

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var dockerBuildArg arrayFlags

type BuildAndDeploy struct{}

func (BuildAndDeploy) Run() error {
	folder, _ := c.ConfigFetch("docker.images.folder", ".")
	dockerRepo := c.RequiredConfigFetch("docker.images.dockerRepo")

	c.LogInfo("Build %s from folder %s", dockerRepo, folder)

	// dockerUser, _ := c.ConfigFetch("docker.user")

	gitSha, _ := c.CaptureCommand("git", "rev-parse", "HEAD")
	dockerTag, _ := c.ConfigFetch("docker.tag", strings.TrimSpace(string(gitSha)))

	dockerBuildCommand := []string{"docker", "build", "-t", dockerRepo, folder}

	if args, ok := c.ConfigFetch("docker.image.build_args"); ok {
		dockerBuildCommand = append(dockerBuildCommand, strings.Split(args, " ")...)
	}

	err := c.Command(dockerBuildCommand...)
	if err != nil {
		return err
	}

	dockerRepoWithTag := fmt.Sprintf("%s:%s", dockerRepo, dockerTag)
	// c.Command("docker", "login", "-u", dockerUser, "-p", "$DOCKER_PASS")
	err = c.Command("docker", "tag", dockerRepo, dockerRepoWithTag)
	if err != nil {
		return err
	}
	err = c.Command("docker", "push", dockerRepo)
	if err != nil {
		return err
	}
	err = c.Command("docker", "push", dockerRepoWithTag)
	if err != nil {
		return err
	}

	return nil
}
