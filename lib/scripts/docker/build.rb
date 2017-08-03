module Docker
  class Build
    class << self
      def description
        <<~MARKDOWN
          Uses docker to build the docker image for the current project.
        MARKDOWN
      end

      def bin_dependencies
        ["docker"]
      end
    end

    def run
      required_env("DOCKER_IMAGE")

      # set image tag if it hasnt been set
      # Default: git tag
      # Description: My favoruite thing
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
