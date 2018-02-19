package CIScriptsHelpers

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
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

func init() {
	var configFile string
	var ok bool
	if configFile, ok = os.LookupEnv("CI_SCRIPTS_CONFIG"); !ok {
		configFile = "./.ci_scripts"
	}

	viper.SetConfigName(path.Base(configFile)) // no need to include file extension
	viper.AddConfigPath(path.Dir(configFile))  // set the path of your config file

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		LogWarning("No configuration file loaded\n")
	}
}

// ConfigCheck sets environment variable key to value if it is not already set. Used to set default env values
func ConfigSetDefault(key, value string) {
	viper.SetDefault(key, value)
}

// RequiredEnv ensures an environment variable is set. If it is not set, it will exit with status code 1
func RequiredConfigKey(key string) {
	if s := viper.GetString(key); s == "" {
		PrettyExit(1, "Required environment variable %s not set", key)
	}
}

// EnvFetch fetchs an environment variables. Sets the value to the default value if it does not exist.
func ConfigFetch(key string, defaultValue ...string) (value string, ok bool) {
	value = viper.GetString(key)
	ok = (value != "")

	if !ok {
		if len(defaultValue) > 0 {
			value = defaultValue[0]
			ok = true
		}
	}

	return value, ok
}
