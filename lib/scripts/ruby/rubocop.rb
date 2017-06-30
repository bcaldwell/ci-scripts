module Ruby
  extend self
  def rubocop
    unless installed?("rubocop")
      command("gem", "install", "rubocop")
    end
    # ugh check bundle rubocop as well
    # puts test_command?("bundle exec rubocop")

    command("rubocop")
  end
end
