module Docker
  class PushBranch
    def run
      env_require("DOCKER_IMAGE")

      # set image tag if it hasnt been set
      # Default: current git hash
      env_check("IMAGE_TAG", `git rev-parse HEAD`.strip)

      branch = `git rev-parse --abbrev-ref HEAD`.strip

      # tag image
      command("docker tag \"$DOCKER_IMAGE:$IMAGE_TAG\" \"$DOCKER_IMAGE:#{branch}\"")

      # push docker image
      command("docker push \"$DOCKER_IMAGE:#{branch}\"")
    end
  end
end
