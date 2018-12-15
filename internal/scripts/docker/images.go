package CIScriptsDocker

import (
	"fmt"
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
	"os"
	"strings"
)

type BuildAndPushImage struct {
	Folder     string
	DockerRepo string
	DockerUser string
	DockerTags []string
}

func (b *BuildAndPushImage) Run() error {
	b.Folder, _ = c.ConfigFetch("docker.images.folder", ".")
	b.DockerRepo = c.RequiredConfigFetch("docker.images.dockerRepo")

	c.LogInfo("Build %s from folder %s", b.DockerRepo, b.Folder)

	b.DockerUser, _ = c.ConfigFetch("docker.user")

	dockerTagsString, _ := c.ConfigFetch("docker.tags", "_tags, _sha, latest")

	b.DockerTags = strings.Split(dockerTagsString, ",")
	b.parseTags()

	fmt.Println(b.DockerTags)

	dockerBuildCommand := []string{"docker", "build", "-t", b.DockerRepo, b.Folder}

	if args, ok := c.ConfigFetch("docker.image.build_args"); ok {
		dockerBuildCommand = append(dockerBuildCommand, strings.Split(args, " ")...)
	}

	err := c.Command(dockerBuildCommand...)
	if err != nil {
		return err
	}

	if b.DockerUser != "" && os.Getenv("DOCKER_PASS") != "" {
		c.Command("sh", "-c", fmt.Sprintf("docker login -u %s -p $DOCKER_PASS", b.DockerUser))
	}

	for _, tag := range b.DockerTags {
		dockerRepoWithTag := fmt.Sprintf("%s:%s", b.DockerRepo, tag)
		err = c.Command("docker", "tag", b.DockerRepo, dockerRepoWithTag)
		if err != nil {
			return err
		}
		err = c.Command("docker", "push", dockerRepoWithTag)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *BuildAndPushImage) parseTags() {
	tags := []string{}
	for _, tag := range b.DockerTags {
		switch strings.TrimSpace(tag) {
		case "_tags":
			tagList, _ := c.CaptureCommand("git", "tag", "-l", "--points-at", "HEAD")
			for _, tag := range strings.Split(string(tagList), "\n") {
				if tag != "" {
					tags = append(tags, strings.TrimSpace(tag))
				}
			}
		case "_sha":
			tag, _ := c.CaptureCommand("git", "rev-parse", "HEAD")
			tags = append(tags, strings.TrimSpace(string(tag)))
		default:
			tags = append(tags, strings.TrimSpace(tag))
		}
	}
	b.DockerTags = tags
}
