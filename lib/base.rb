##!/usr/bin/ruby

BASE_SCRIPT_PATH = "https://raw.githubusercontent.com/benjamincaldwell/ci-scripts/master/"

def command(*options)
  puts("\x1b[34m#{options.join " "}\x1b[0m")
  t = Time.now
  system(*options)
  puts("\x1b[32m#{(Time.now - t).round(2)}s\x1b[0m \n ")
end

def env_check(key, value)
  unless ENV[key]
    puts "Setting #{key} to #{value}"
    ENV[key] = value
  end
end

def required_env(key)
  unless ENV[key]
    puts "\x1b[31mRequired environment variable #{key} not set\x1b[0m"
    exit 1
  end
end

def from_env(key, default="")
  if ENV[key]
    ENV[key]
  else
    default
  end
end

def run_script(path)
  require 'net/http' unless defined? Net::HTTP
  uri = URI.join(BASE_SCRIPT_PATH, path)
  res = Net::HTTP.get_response(uri)

  unless res.code == "200"
    puts "Unable to fetch #{path}\n"
    exit 1
  end

  eval(res.body)
end
