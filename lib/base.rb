# #!/usr/bin/ruby

BASE_SCRIPT_PATH = "https://raw.githubusercontent.com/benjamincaldwell/ci-scripts/master/".freeze

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
  log_info(options.join(" "))
  t = Time.now
  system(*options)
  log_success("#{(Time.now - t).round(2)}s\n ")
  exit $CHILD_STATUS.exitstatus if $CHILD_STATUS.exitstatus != 0
end

def timed_run(name)
  log_info(name)
  t = Time.now
  yield
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

def env_fetch(key, default = "")
  if ENV[key]
    ENV[key]
  else
    default
  end
end

def classify(s)
  s = s.to_s.split('_').collect(&:capitalize).join
  s[0] = s[0].capitalize
  s
end

def run_script(script_name)
  require script_name

  script_parts = script_name.split("/")
  function_name = script_parts.pop
  module_name = ""

  script_parts.each do |part|
    module_name += "::" unless module_name.empty?
    module_name += classify(part)
  end

  eval("#{module_name}.#{function_name}")
end
