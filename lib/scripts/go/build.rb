module Go
  class Build
    class << self
      def description
        <<-MARKDOWN
          Builds go binary using gox which enables parallel builds
        MARKDOWN
      end
    end

    def run
      command("go", "get", "github.com/mitchellh/gox")

      build_command = ["gox"]

      build_command = add_option("ldflags", build_command)

      parallel = env_fetch("GO_BUILD_PARALLEL", "2")
      build_command.push("-parallel")
      build_command.push(parallel)

      build_command = add_option("osarch", build_command)
      build_command = add_option("output", build_command)

      if gox_args = env_fetch("GO_BUILD_ARGS")
        build_command.concat(gox_args.split)
      end

      command(*build_command)
    end

    private

    def add_option(name, build_command)
      env_name = "GO_BUILD_#{name.upcase}"
      if val = env_fetch(env_name)
        build_command.push("-#{name}")
        build_command.push("\"#{val}\"")
      end
      build_command
    end
  end
end
