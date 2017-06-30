module Docker
  extend self
  def herokuish
    # set image tag if it hasnt been set
    # also has to support old ruby versions
    dockerfile_contents = <<-DOCKERFILE
      FROM gliderlabs/herokuish

      COPY . /app

      RUN /bin/herokuish buildpack build

      CMD ["/start", "web"]
    DOCKERFILE

    dockerfile_contents = unindent(dockerfile_contents)

    timed_run "Creating herokuish dockerfile" do
      File.write("Dockerfile.herokuish", dockerfile_contents)

      ENV["BUILD_DOCKERFILE"] = "Dockerfile.herokuish"
    end

    run_script("docker/build")
  end
end
