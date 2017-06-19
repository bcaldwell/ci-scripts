##!/usr/bin/ruby

BASE_SCRIPT_PATH = "https://raw.githubusercontent.com/benjamincaldwell/ci-scripts/master/"

def log_info(s)
  puts("\x1b[34m#{s}\x1b[0m")
end

def log_success(s)
  puts("\x1b[32m#{s}\x1b[0m")
end

def log_error(s)
  puts("\x1b[31m#{s}\x1b[0m")
end

def command(*options)
  log_info(options.join " ")
  t = Time.now
  system(*options)
  log_success("#{(Time.now - t).round(2)}s\n ")
end

def env_check(key, value)
  unless ENV[key]
    puts "Setting #{key} to #{value}"
    ENV[key] = value
  end
end

def required_env(key)
  unless ENV[key]
    log_error "Required environment variable #{key} not set"
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
