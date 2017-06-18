##!/usr/bin/ruby

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

def run_script(path)
  raise "Not implemented"
end
