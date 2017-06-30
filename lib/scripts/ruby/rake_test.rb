module Ruby
  extend self
  def rake_test
    if test_command?("bundler", "exec", "rake", "-V")
      return command?("bundler", "exec", "rake", "test")
    end

    unless installed?("rake")
      command("gem", "install", "rake")
    end
    command("rake", "test")
  end
end
