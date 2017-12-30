module Docker
  class Login
    def run
      # set image tag if it hasnt been set
      env_require("DOCKER_USERNAME")
      env_require("DOCKER_PASSWORD")

      env_check("DOCKER_EMAIL", "ci@ci-runner.com")

      # Default: hub.docker.com
      docker_registry = env_fetch("DOCKER_REGISTRY", nil)

      # login to docker hub
      login_command = 'docker login -u "$DOCKER_USERNAME" -p "$DOCKER_PASSWORD" -e "$DOCKER_EMAIL"'
      login_command += " #{docker_registry}" if docker_registry
      command(login_command)
    end
  end
end
