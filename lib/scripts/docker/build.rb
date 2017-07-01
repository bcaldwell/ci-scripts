module Docker
  class Build
    def run
      required_env("DOCKER_IMAGE")

      # set image tag if it hasnt been set
      env_check("IMAGE_TAG", `git rev-parse HEAD`.strip)

      dockerfile = env_fetch("BUILD_DOCKERFILE", "Dockerfile")
      # project_folder = env_fetch("DOCKER_FOLDER", ".")

      # build docker image
      command("docker build --pull -t \"$DOCKER_IMAGE:$IMAGE_TAG\" -f #{dockerfile} .")

      # push docker image
      command('docker push "$DOCKER_IMAGE:$IMAGE_TAG"')
    end
  end
end
