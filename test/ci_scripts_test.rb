require "test_helper"
require "pathname"

class CIScriptsTest < Minitest::Test
  def test_that_it_has_a_version_number
    refute_nil CIScripts::VERSION
  end

  def test_that_run_script_returns_false_if_script_doesnt_exist
    refute CIScripts::Script.new("sdfgdfhjghgdfv").run
  end

  def test_that_run_script_runs_a_script
    assert CIScripts::Script.new("demo/test").run
  end

  def test_that_all_script_use_the_correct_format
    script_path = Pathname.new(File.expand_path("./lib/scripts"))
    Dir.glob('./lib/scripts/**/*.rb') do |script_name|
      absolute_path = Pathname.new(File.expand_path(script_name))
      script_name = absolute_path.relative_path_from(script_path).to_s.split(".")[0]

      script = CIScripts::Script.new(script_name)

      # check class name
      class_name = script.instance_variable_get(:@class_name)
      klass = Object.const_get(class_name)

      # check for run function
      assert klass.method_defined? :run
    end
  end
end
