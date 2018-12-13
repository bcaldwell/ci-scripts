package CIScriptsHelpers

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"regexp"
)

func init() {
	var configFile string
	var ok bool
	if configFile, ok = os.LookupEnv("CI_SCRIPTS_CONFIG"); !ok {
		configFile = "./.ci_scripts"
	}
	// remove file extension
	configFile = strings.TrimSuffix(configFile, filepath.Ext(configFile))

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

func RequiredConfigFetch(key string) (value string) {
	if s, ok := ConfigFetch(key); !ok {
		PrettyExit(1, "Required configuration variable %s not set", key)
	} else {
		return s
	}
	return ""
}

func ParseCLIArguments() {
	flagParser, _ := regexp.Compile("--?([^=]+)=?(.+)?")

	flagName := ""

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			match := flagParser.FindStringSubmatch(arg)
			if len(match) < 1 {
				LogError("Failed to parse flag with regex --?([^=]+)=?(.+): ", arg)
			}
			flagName = strings.Replace(match[1], "-", ".", -1)

			if len(match) == 3 && match[2] != "" {
				viper.Set(flagName, match[2])
				flagName = ""
			}
		} else if flagName != "" {
			viper.Set(flagName, arg)
			flagName = ""
		}
	}
}
