module Docker
  class PushLatest
    def run
      latest_branch = env_fetch("DOCKER_LATEST_BRANCH", "master")

      branch = `git rev-parse --abbrev-ref HEAD`.strip
      return if branch != latest_branch

      required_env("DOCKER_IMAGE")

      # set image tag if it hasnt been set
      # Default: current git hash
      env_check("IMAGE_TAG", `git rev-parse HEAD`.strip)

      # tag latest image
      command('docker tag "$DOCKER_IMAGE:$IMAGE_TAG" "$DOCKER_IMAGE:latest"')

      # push docker image
      command('docker push "$DOCKER_IMAGE:latest"')
    end
  end
end
