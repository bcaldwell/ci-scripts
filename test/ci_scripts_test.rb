require "test_helper"

class CIScriptsTest < Minitest::Test
  def test_that_it_has_a_version_number
    refute_nil CIScripts::VERSION
  end

  def test_that_run_script_returns_false_if_script_doesnt_exist
    refute CIScripts.run_script("sdfgdfhjghgdfv")
  end

  def test_that_run_script_runs_a_script
    assert CIScripts.run_script("demo/test")
  end
end
