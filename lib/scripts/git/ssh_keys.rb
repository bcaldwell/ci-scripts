module Git
  extend self
  def ssh_keys
    timed_run "something" do
      puts "hello"
    end
  end
end
