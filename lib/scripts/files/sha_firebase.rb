require 'firebase'
require 'digest'

module Files
  class ShaFirebase
    class << self
      def description
        <<-MARKDOWN
          Calculates the sha of specified files and uploads the result to firebase
        MARKDOWN
      end
    end

    def run
      firebase_uri = env_require("FIREBASE_URI")
      firebase_api_key = env_require("FIREBASE_API_KEY")

      sha_version = env_require("SHA_VERSION")
      sha_folder = env_require("SHA_FOLDER")

      shas = {
        created: Firebase::ServerValue::TIMESTAMP
      }

      timed_run "Generating SHA256" do
        sha256 = Digest::SHA256.new

        return puts "#{sha_folder} does not exist" unless File.exists?(sha_folder)

        Dir.foreach(sha_folder) do |file|
          next if File.directory?(file)
          escaped_filename = firebase_escape(file.to_s)
          shas[escaped_filename.to_sym] = sha256.file(file).hexdigest
        end
      end

      timed_run "Uploading to firebase" do
        firebase = Firebase::Client.new(firebase_uri, firebase_api_key)

        url = File.join(git_url, sha_version)
        latest_url = File.join(git_url, "latest")
        url = firebase_escape(url)
        latest_url = firebase_escape(latest_url)

        response = firebase.set(url, shas)
        unless response.code == 200
          log_error("Failed to upload to Firebase")
          nice_exit(0, response.body)
        end

        response = firebase.set(latest_url, {
          version: sha_version,
        }.merge!(shas))
        unless response.code == 200
          log_error("Failed to upload to Firebase")
          nice_exit(0, response.body)
        end
      end
    end

    private

    def git_url
      git_remotes = capture_command("git", "remote", "-v")
      url = /(?:git@|https:\/\/)([^\s]+)/.match(git_remotes)[1]
      url.gsub(":", "/")
    end

    def firebase_escape(s)
      s.gsub(".", "_")
    end
  end
end
