module Files
  class Tarball
    def run
      # env_require("files.tarball.output_directory")
      tar_output_dir = env_require("FILES_TARBALL_OUTPUT_DIRECTORY")
      # require "YAML"
      # thing = YAML.load_file("ci_scripts.yml")

      # get files to add to tarball
      # Description: comma separated for
      tar_files = env_require("FILES_TARBALL_FILES").split(",")
      tar_included = env_fetch("FILES")
    end
  end
end
