package CIScriptsDocker

import (
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
	"github.com/bcaldwell/sshtun"
	"github.com/kevinburke/ssh_config"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// docker/build_and_deploy folder --copy-dockerfile-to-root --build-arg (arg passed to docker build)

type BuildAndDeploy struct{}

func (BuildAndDeploy) Run() error {
	folder, _ := c.ConfigFetch("docker.swarm.folder", ".")
	deployStack, _ := c.ConfigFetch("docker.swarm.stack", folder)
	deployFile := c.RequiredConfigFetch("docker.swarm.deployfile")
	host := c.RequiredConfigFetch("docker.swarm.host")
	dockerSock := path.Join(os.TempDir(), "/docker.sock")

	// probably need to check if its a full path first...
	deployFile = path.Join(folder, deployFile)

	if _, err := os.Stat(dockerSock); !os.IsNotExist(err) {
		// path/to/whatever exists
		os.Remove(dockerSock)
	}

	f, _ := os.Open(path.Join(os.Getenv("HOME"), ".ssh", "config"))
	sshConfig, _ := ssh_config.Decode(f)

	sshTun := sshtun.NewUnix(dockerSock, host, "/var/run/docker.sock")

	if user, ok := c.ConfigFetch("docker.swarm.user"); ok {
		sshTun.SetUser(user)
	} else {
		user, _ = sshConfig.Get(host, "User")
		if user != "" {
			sshTun.SetUser(user)
		}
	}

	if keyfile, ok := c.ConfigFetch("docker.swarm.keyfile"); ok {
		sshTun.SetKeyFile(keyfile)
	} else {
		keyfile, _ = sshConfig.Get(host, "IdentityFile")
		if keyfile != "" {
			sshTun.SetKeyFile(keyfile)
		} else {
			// use ssh key if there is only 1 in ~/.ssh folder
			keys, _ := filepath.Glob(path.Join(os.Getenv("HOME"), ".ssh", "id_*"))
			keys = c.Filter(keys, func(s string) bool {
				return !strings.HasSuffix(s, ".pub")
			})

			if len(keys) == 1 {
				sshTun.SetKeyFile(keys[0])
			}
		}
	}

	go func() {
		err := sshTun.Start()
		if err != nil {
			c.LogError(err.Error())
		}
	}()

	//         docker -H localhost:2374 stack deploy --compose-file deploy/jupyterhub.yml jupyterhub

	c.Command("docker", "-H", "unix://"+dockerSock, "stack", "deploy", "--compose-file", deployFile, deployStack)
	sshTun.Stop()

	return nil
}
