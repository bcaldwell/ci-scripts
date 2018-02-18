package CIScriptsHelpers

import (
	"os"
)

// # env helpers
// def env_check(key, value)
//   unless ENV[key]
//     puts "Setting #{key} to #{value}"
//     ENV[key] = value
//   end
// end

// def required_env(key)
//   unless ENV[key]
//     log_error "Required environment variable #{key} not set"
//     exit 1
//   end
// end

// def env_fetch(key, default = "")
//   if ENV[key]
//     ENV[key]
//   else
//     default
//   end
// end

// EnvCheck sets environment variable key to value if it is not already set. Used to set default env values
func EnvCheck(key, value string) {
	if _, ok := os.LookupEnv(key); !ok {
		os.Setenv(key, value)
	}
}

// RequiredEnv ensures an environment variable is set. If it is not set, it will exit with status code 1
func RequiredEnv(key string) {
	if _, ok := os.LookupEnv(key); !ok {
		// LogError("Required environment variable %s not set", key)
		os.Exit(1)
	}
}

// EnvFetch fetchs an environment variables. Sets the value to the default value if it does not exist.
func EnvFetch(key string, defaultValue ...string) (value string, ok bool) {
	if value, ok = os.LookupEnv(key); !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
			ok = true
		}
	}

	return value, ok
}
