# coding: utf-8
lib = File.expand_path("../lib", __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require "ci_scripts/version"

Gem::Specification.new do |spec|
  spec.name          = "ci-scripts"
  spec.version       = CIScripts::VERSION
  spec.authors       = ["Benjamin Caldwell"]
  spec.email         = ["caldwellbenjamin8@gmail.com"]

  spec.summary       = "A collection fo scripts for commomly run scripts in CI."
  # spec.description   = %q{TODO: Write a longer description or delete this line.}
  spec.homepage      = "http://github.com/benjamincaldwell/ci-scripts"
  spec.license       = "MIT"

  spec.files         = `git ls-files -z`.split("\x0").reject do |f|
    f.match(%r{^(test|spec|features)/})
  end
  spec.bindir        = "bin"
  spec.executables   = ["ci-scripts"]
  spec.require_paths = ["lib"]

  spec.add_dependency "firebase", "~> 0.2.6"

  spec.add_development_dependency "bundler", "~> 1.15"
  spec.add_development_dependency "rake", "~> 10.0"
  spec.add_development_dependency "minitest", "~> 5.0"
  spec.add_development_dependency "rubocop", "~> 0.49.0"
  spec.add_development_dependency "byebug"
end
