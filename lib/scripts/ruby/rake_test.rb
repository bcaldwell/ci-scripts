module Ruby
  class RakeTest
    class << self
      def description
        <<-MARKDOWN
          Runs ruby tests by executing the `rake test` command. Uses bundler if bundler is installed.
        MARKDOWN
      end

      # def bin_dependencies
      #   ["bundler"]
      # end
    end

    def run
      if test_command?("bundler", "exec", "rake", "-V")
        return command("bundler", "exec", "rake", "test")
      end

      unless installed?("rake")
        command("gem", "install", "rake")
      end
      command("rake", "test")
    end
  end
end
