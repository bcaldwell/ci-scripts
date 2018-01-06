#!/usr/bin/bash
#!/usr/bin/env bash

# set -e
# echo "" > coverage.txt

# for d in $(go list ./... | grep -v vendor); do
#     go test -v -race -coverprofile=profile.out -covermode=atomic $d
#     if [ -f profile.out ]; then
#         cat profile.out >> coverage.txt
#         rm profile.out
#     fi
# done

require 'tempfile'

module Go
  class Test
    class << self
      def description
        <<-MARKDOWN
          Run all test from all packages. Creates coverage.txt which concats all coverprofiles from go test.
        MARKDOWN
      end
    end

    def run
      go_test_package = env_fetch("GO_TEST_PACKAGE", ".")
      go_coverage_output = env_fetch("GO_COVERAGE_OUTPUT", "./coverage.txt")

      output = capture_command("sh", "-c", "go list #{go_test_package}/... | grep -v vendor")
      packages = output.split("\n")

      cover_output = ""

      packages.each do |package|
        output_file = Tempfile.new('go test')

        command("go", "test", "-v", "-race", "-coverprofile=#{output_file.path}", "-covermode=atomic", package)

        cover_output += output_file.read
        output_file.close
        output_file.unlink
      end

      File.write(go_coverage_output, cover_output)
    end
  end
end
