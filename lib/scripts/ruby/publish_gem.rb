module Ruby
  class PublishGem
    def run
      # Bundler release scripts
      # https://github.com/bundler/bundler/blob/master/lib/bundler/gem_helper.rb
      # also gemspec loading https://github.com/bundler/bundler/blob/ff4a522e8e75eb4ce5675a99698fb3df23b680be/lib/bundler.rb#L406
      out, stat = capture_command("bundle", "exec", "rake", "-T")

      if stat && out.include?("rake release[remote]")
        command("bundle", "exec", "rake", "release")
      else
        log_error("bundler/gem_tasks is required to use this script")
      end
    end
  end
end
