package CIScriptsDocker

import (
	"fmt"
	"os"
	"strings"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type DockerBase struct {
	DockerRepo string
	DockerUser string
	DockerTags []string
}

type BuildAndPushImage struct {
	DockerBase
	Folder string
}

type CombineAndPushImage struct {
	DockerBase
	AmendTags []string
}

func (b *BuildAndPushImage) Run() error {
	b.Folder, _ = c.ConfigFetch("docker.images.folder", ".")
	b.DockerBase.setup()

	c.LogInfo("Build %s from folder %s", b.DockerRepo, b.Folder)
	dockerContextName := strings.Replace(b.DockerRepo, "/", "-", -1)

	err := c.Command("docker", "context", "create", dockerContextName)
	if err != nil {
		return err
	}

	err = c.Command("docker", "buildx", "create", dockerContextName, "--use")
	if err != nil {
		return err
	}

	err = c.Command("docker", "run", "--privileged", "--rm", "tonistiigi/binfmt", "--install", "arm64,arm")
	if err != nil {
		return err
	}

	// docker buildx build . --platform=linux/arm64,linux/amd64 --tag my/image:0.1 --tag my/image:latest --pull --push --no-cache
	dockerBuildCommand := []string{"docker", "buildx", "build", b.Folder, "--progress", "plain", "--platform", c.ConfigFetchWithDefault("docker.image.platform", "linux/amd64,linux/arm64"), "--push"}

	if args, ok := c.ConfigFetch("docker.image.build_args"); ok {
		dockerBuildCommand = append(dockerBuildCommand, strings.Split(args, " ")...)
	}

	for _, tag := range b.DockerTags {
		dockerRepoWithTag := fmt.Sprintf("%s:%s", b.DockerRepo, tag)
		dockerBuildCommand = append(dockerBuildCommand, "--tag", dockerRepoWithTag)
	}

	err = c.Command(dockerBuildCommand...)
	if err != nil {
		return err
	}

	return nil
}

func (b *CombineAndPushImage) Run() error {
	b.AmendTags = strings.Split(c.RequiredConfigFetch("docker.combine.amend_tags"), ",")
	b.DockerBase.setup()

	dockerManifestCommand := []string{"docker", "manifest", "create", b.DockerRepo}
	for _, tag := range b.AmendTags {
		dockerManifestCommand = append(dockerManifestCommand, "--amend", fmt.Sprintf("%s:%s", b.DockerRepo, tag))
	}

	err := c.Command(dockerManifestCommand...)
	if err != nil {
		return err
	}

	for _, tag := range b.DockerTags {
		taggedImage := fmt.Sprintf("%s:%s", b.DockerRepo, tag)
		err = c.Command("docker", "manifest", "tag", b.DockerRepo, taggedImage)
		if err != nil {
			return err
		}
		err = c.Command("docker", "manifest", "push", taggedImage)
	}

	return nil
}

func (b *DockerBase) setup() {
	b.DockerRepo = c.RequiredConfigFetch("docker.images.dockerRepo")

	b.DockerUser, _ = c.ConfigFetch("docker.user")

	dockerTagsString, _ := c.ConfigFetch("docker.tags", "_tags, _sha, latest")

	b.DockerTags = strings.Split(dockerTagsString, ",")
	b.parseTags()

	fmt.Println(b.DockerTags)

	if b.DockerUser != "" && os.Getenv("DOCKER_PASS") != "" {
		c.Command("sh", "-c", fmt.Sprintf("docker login -u %s -p $DOCKER_PASS", b.DockerUser))
	}
}

func (b *DockerBase) parseTags() {
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
