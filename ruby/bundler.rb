module Ruby
  extend self
  def bundler
    unless installed?("bundler")
      command("gem", "install", "bundler", "--no-ri", "--no-rdoc")
    end

    unless test_command?("bundler", "check")
      command("bundler", "install")
    end
  end
end

Ruby.bundler if __FILE__ == $PROGRAM_NAME
