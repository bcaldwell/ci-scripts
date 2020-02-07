package kubernetes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
	"github.com/bcaldwell/ci-scripts/pkg/ejsonsecret"
	"github.com/kevinburke/ssh_config"
	"github.com/rgzr/sshtun"
)

const (
	KUBECONFIG_PATH = "~/.kube/config"
)

type kubeNamespace struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   map[string]string `json:"metadata"`
}

type Deploy struct {
	kubeconfigPath string
}

func (d *Deploy) Run() error {
	// github repo to clone to access the config
	// configRepo, _ := c.ConfigFetch("kubernetes.deploy.repo")
	// path to the deployment config folder: relative to the github repo root or absolte if local
	configFolder := c.RequiredConfigFetch("kubernetes.deploy.configFolder")
	// helm chart configuration folder
	helmValuesFolder, _ := c.ConfigFetch("kubernetes.deploy.folder.helmvalues", "helmvalues")
	// folder with ejson secrets to deploy
	secretsFolder, _ := c.ConfigFetch("kubernetes.deploy.folder.secrets", "secrets")
	// folder with pre deployment files to deploy
	predeployFolder, _ := c.ConfigFetch("kubernetes.deploy.folder.predeploy", "predeploy")
	// bastion host to port forward kubernetes api server from
	bastionHost, _ := c.ConfigFetch("kubernetes.deploy.bastion")
	// namespace to deploy everything to
	namespace := c.RequiredConfigFetch("kubernetes.deploy.namespace")
	// env varaible to get kubeconfig from
	kubeconfigEnv, _ := c.ConfigFetch("kubernetes.deploy.kubeconfigEnv", "KUBE_CONFIG")
	// kube config path
	kubeconfigPath, _ := c.ConfigFetch("kubernetes.deploy.kubeconfig", KUBECONFIG_PATH)
	// helm chart path
	helmChart := c.RequiredConfigFetch("kubernetes.deploy.helmChart")
	// helm release name
	releaseName := c.RequiredConfigFetch("kubernetes.deploy.releaseName")

	predeployFolder = d.resolveFolderFromConfig(configFolder, predeployFolder)
	secretsFolder = d.resolveFolderFromConfig(configFolder, secretsFolder)
	helmValuesFolder = d.resolveFolderFromConfig(configFolder, helmValuesFolder)

	if bastionHost != "" {
		c.LogInfo("Setting up bastion host tunnel")

		sshTun := d.setupPortForward(bastionHost, 6443, 6443)
		defer sshTun.Stop()
	}

	kubeconfig, err := d.getKubeConfig(kubeconfigPath, kubeconfigEnv)
	if err != nil {
		return err
	}

	fmt.Println(kubeconfig)

	os.Setenv("KUBECONFIG", kubeconfig)

	err = d.createNamespace(namespace)
	if err != nil {
		return fmt.Errorf("error while creating namespace %s %w", namespace, err)
	}

	if predeployFolder != "" {
		err = d.kubectlDeployFolder(predeployFolder, namespace)
		if err != nil {
			return fmt.Errorf("failed to deploy predeploy folder %s %w", predeployFolder, err)
		}
	}

	if secretsFolder != "" {
		err = d.deploySecretsFolder(secretsFolder)
		if err != nil {
			return fmt.Errorf("failed to deploy secrets folder %s %w", secretsFolder, err)
		}
	}

	err = d.deployHelm(releaseName, helmChart, namespace, helmValuesFolder)
	if err != nil {
		return fmt.Errorf("failed to deploy helm chart %s %w", helmChart, err)
	}

	return nil
}

func (*Deploy) setupPortForward(host string, localPort, remotePort int) *sshtun.SSHTun {
	f, _ := os.Open(path.Join(os.Getenv("HOME"), ".ssh", "config"))
	sshConfig, _ := ssh_config.Decode(f)

	sshTun := sshtun.New(localPort, host, remotePort)

	if user, ok := c.ConfigFetch("docker.kubernetes.bastionUser"); ok {
		sshTun.SetUser(user)
	} else {
		user, _ = sshConfig.Get(host, "User")
		if user != "" {
			sshTun.SetUser(user)
		}
	}

	if keyfile, ok := c.ConfigFetch("docker.kubernetes.bastionKeyfile"); ok {
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

	return sshTun
}

func (*Deploy) getKubeConfig(kubeconfigPath string, kubeconfigEnv string) (string, error) {
	kubeconfigPath = c.ExpandPath(kubeconfigPath)

	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		if _, err := os.Stat(kubeconfig); err == nil {
			c.LogInfo("Using kube config defined in KUBECONFIG environment variable %s", kubeconfig)
			return kubeconfig, nil
		}
	}

	if _, err := os.Stat(kubeconfigPath); err == nil {
		c.LogInfo("Using existing kube config found at %s", kubeconfigPath)
	} else if kubeconfig := os.Getenv(kubeconfigEnv); kubeconfig != "" && os.IsNotExist(err) {
		c.LogInfo("Creating kubeconfig from environment variable %s", kubeconfigEnv)
		decodedKubeconfig, err := base64.StdEncoding.DecodeString(kubeconfig)
		if err != nil {
			return "", fmt.Errorf("Failed to decode base64 encode kube config %w", err)
		}

		f, err := os.Create(kubeconfigPath)
		if err != nil {
			return "", fmt.Errorf("Failed to creating kube config from environemt %w", err)
		}
		defer f.Close()

		if _, err := f.Write(decodedKubeconfig); err != nil {
			return "", fmt.Errorf("Failed to writing kube config from environemt %w", err)
		}
		if err := f.Sync(); err != nil {
			return "", fmt.Errorf("Failed to syncing kube config from environemt %w", err)
		}
	} else {
		return "", fmt.Errorf("Failed to find kubeconfig to use. Either set KUBE_CONFIG environment variable or kubeconfig config")
	}

	return kubeconfigPath, nil
}

func (*Deploy) createNamespace(namespace string) error {
	ns := kubeNamespace{
		APIVersion: "v1",
		Kind:       "Namespace",
		Metadata: map[string]string{
			"name": namespace,
		},
	}

	manifest, err := json.Marshal(ns)
	if err != nil {
		return fmt.Errorf("failed to marshal kubernetes namespace %w", err)
	}

	c.LogInfo("creating namespace %s", namespace)

	kubeApply := exec.Command("kubectl", "apply", "-f", "-")
	kubeApply.Stdin = bytes.NewReader(manifest)

	return c.RunCommand(kubeApply)
}

func (*Deploy) kubectlDeployFolder(folder, namespace string) error {
	if _, err := os.Stat(folder); err == nil {
		c.LogInfo("Deploying folder %s", folder)
		// -R for recursive
		return c.Command("kubectl", "apply", "-R", "-f", folder)
	}

	return nil
}

func (*Deploy) deploySecretsFolder(folder string) error {
	if _, err := os.Stat(folder); err == nil {
		c.LogInfo("Deploying secrets folder %s", folder)

		files, err := c.RecursiveFilesInFolder(folder)

		if err != nil {
			return fmt.Errorf("failed to get files in secrets folder %s %w", folder, err)
		}

		for _, f := range files {
			err = ejsonsecret.DeploySecret(f, os.Getenv("EJSON_KEY"))
			if err != nil {
				return fmt.Errorf("failed to deploy secret %s %w", f, err)
			}
		}
	}

	return nil
}

func (*Deploy) deployHelm(releaseName, helmchart, namespace, configFolder string) error {
	c.LogInfo("Deploying helm chart %s with release %s into %s", helmchart, releaseName, namespace)

	helmCmd := []string{"helm", "upgrade", "--install", "-n", namespace, releaseName, helmchart}

	if _, err := os.Stat(configFolder); err == nil {
		files, err := c.RecursiveFilesInFolder(configFolder)

		if err != nil {
			return fmt.Errorf("failed to get files in helm config folder %s %w", configFolder, err)
		}

		for _, f := range files {
			// todo copy these files and expand them
			helmCmd = append(helmCmd, "-f", f)
		}
	}

	return c.Command(helmCmd...)
}

func (*Deploy) resolveFolderFromConfig(configFolder, configPath string) string {
	return path.Join(configFolder, configPath)
}
