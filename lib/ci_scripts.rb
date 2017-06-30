require "ci_scripts/version"

require "ci_scripts/helpers"

# required for using english versions of ruby globals
# https://github.com/CocoaPods/cocoapods-downloader/issues/28
require "English"

module CIScripts
  def self.run_script(script_name)
    script_name = script_name.strip
    full_path = File.join(File.dirname(__FILE__), "scripts", script_name)

    unless File.exist?("#{full_path}.rb")
      log_error "#{script_name} does not exists"
      return false
    end

    require full_path

    script_parts = script_name.split("/")
    function_name = script_parts.pop
    module_name = ""

    script_parts.each do |part|
      module_name += "::" unless module_name.empty?
      module_name += classify(part)
    end

    Object.const_get(module_name).send(function_name)
  end
end
