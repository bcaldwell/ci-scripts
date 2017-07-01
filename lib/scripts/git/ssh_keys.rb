module Git
  class SshKeys
    def run
      timed_run "something" do
        puts "hello"
      end
    end
  end
end
