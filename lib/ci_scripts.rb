require "ci_scripts/version"

require "ci_scripts/helpers"

# required for using english versions of ruby globals
# https://github.com/CocoaPods/cocoapods-downloader/issues/28
require "English"

module CIScripts
  class Script
    def initialize(script_name)
      script_name = script_name.strip
      full_path = File.join(File.dirname(__FILE__), "scripts", script_name)

      unless File.exist?("#{full_path}.rb")
        log_error "#{script_name} does not exists"
        return
      end

      require full_path

      @class_name = parse_script_name(script_name)
    end

    def run
      return false unless @class_name

      result = Object.const_get(@class_name).new.send("run")
      return true if result.nil?
      result
    end

    private

    def parse_script_name(script)
      module_name = ""

      script_parts = script.split("/")
      # function_name = script_parts.pop

      script_parts.each do |part|
        module_name += "::" unless module_name.empty?
        module_name += classify(part)
      end

      module_name
    end
  end
end
