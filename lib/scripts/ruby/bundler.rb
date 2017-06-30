module Ruby
  extend self
  def bundler
    unless installed?("bundler")
      command("gem", "install", "bundler", "--no-ri", "--no-rdoc")
    end

    install_path = env_fetch("BUNDLER_INSTALL_PATH", "vendor")

    unless test_command?("bundler", "check")
      command("bundler", "install", "--path", install_path)
    end
  end
end
