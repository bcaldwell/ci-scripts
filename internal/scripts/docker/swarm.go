package CIScriptsDocker

import (
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
	"github.com/bcaldwell/sshtun"
	"os"
	"path"
)

// docker/build_and_deploy folder --copy-dockerfile-to-root --build-arg (arg passed to docker build)

type BuildAndDeploy struct{}

func (BuildAndDeploy) Run() error {
	folder, _ := c.ConfigFetch("docker.swarm.folder", ".")
	deployFile := c.RequiredConfigFetch("docker.swarm.deployFile")
	masterIP := c.RequiredConfigFetch("docker.swarm.masterIP")

	dockerSock := path.Join(os.TempDir(), "/docker.sock")

	if _, err := os.Stat(dockerSock); !os.IsNotExist(err) {
		// path/to/whatever exists
		os.Remove(dockerSock)
	}

	sshTun := sshtun.NewUnix(dockerSock, masterIP, "/var/run/docker.sock")

	if user, ok := c.ConfigFetch("docker.swarm.user"); ok {
		sshTun.SetUser(user)
	}

	if keyfile, ok := c.ConfigFetch("docker.swarm.keyFile"); ok {
		sshTun.SetKeyFile(keyfile)
	}

	go func() {
		err := sshTun.Start()
		if err != nil {
			c.LogError(err.Error())
		}
	}()

	//         docker -H localhost:2374 stack deploy --compose-file deploy/jupyterhub.yml jupyterhub

	c.Command("docker", "-H", "unix://"+dockerSock, "stack", "deploy", deployFile, folder)
	sshTun.Stop()

	return nil
}
