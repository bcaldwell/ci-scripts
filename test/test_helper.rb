$LOAD_PATH.unshift File.expand_path("../../lib", __FILE__)
require "ci_scripts"

require "minitest/autorun"

# haha surpress output
def puts(*s)
end
