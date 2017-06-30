# #!/usr/bin/ruby

# Logging
def log_info(s)
  puts("\x1b[34m#{s}\x1b[0m")
end

def log_success(s)
  puts("\x1b[32m#{s}\x1b[0m")
end

def log_error(s)
  puts("\x1b[31m#{s}\x1b[0m")
end

# Timed runs
def command(*options)
  log_info(options.join(" "))
  t = Time.now
  system(*options)
  log_success("#{(Time.now - t).round(2)}s\n ")
  exit $CHILD_STATUS if $CHILD_STATUS != 0
end

def timed_run(name)
  log_info(name)
  t = Time.now
  yield
  log_success("#{(Time.now - t).round(2)}s\n ")
end

# system helpers
def test_command?(*options)
  system(*options, out: File::NULL)
  $CHILD_STATUS == 0
end

def installed?(binary)
  test_command?("command", "-v", binary)
end

# env helpers
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

# helpers
def classify(s)
  s = s.to_s.split('_').collect(&:capitalize).join
  s[0] = s[0].capitalize
  s
end

def unindent(s)
  indent = s.split("\n").select { |line| !line.strip.empty? }.map { |line| line.index(/[^\s]/) }.compact.min || 0
  s.gsub(/^[[:blank:]]{#{indent}}/, '')
end
