module Git
  extend self
  def ssh_keys
    timed_run "something" do
      puts "hello"
    end
  end
end

Git.ssh_keys if __FILE__ == $PROGRAM_NAME
