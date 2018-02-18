package CIScriptsHelpers

import (
	"fmt"
	"os"
)

// # Logging
// def log_info(s)
//   puts("\x1b[34m#{s}\x1b[0m")
// end

// def log_success(s)
//   puts("\x1b[32m#{s}\x1b[0m")
// end

// def log_error(s)
//   puts("\x1b[31m#{s}\x1b[0m")
// end

// def nice_exit(code, msg = "Something happened")
//   log_error(msg)
//   exit code
// end

// LogInfo logs log line to terminal in blue
func LogInfo(s string, opt ...interface{}) {
	fmt.Printf("\x1b[34m"+s+"\x1b[0m\n", opt...)
}

// LogInfo logs log line to terminal in yellow
func LogWarning(s string, opt ...interface{}) {
	fmt.Printf("\x1b[33m"+s+"\x1b[0m\n", opt...)
}

// LogInfo logs log line to terminal in green
func LogSuccess(s string, opt ...interface{}) {
	fmt.Printf("\x1b[32m"+s+"\x1b[0m\n", opt...)
}

// LogInfo logs log line to terminal in red
func LogError(s string, opt ...interface{}) {
	fmt.Printf("\x1b[31m"+s+"\x1b[0m\n", opt...)
}

func PrettyExit(code int, s string, opt ...interface{}) {
	LogError(s, opt...)
	os.Exit(code)
}
