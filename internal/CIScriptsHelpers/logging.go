package CIScriptsHelpers

import (
	"fmt"
	"os"
	"time"
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
	fmt.Printf(timeWrap("\x1b[34m"+s+"\x1b[0m\n"), opt...)
	// fmt.Printf(timeWrap(s+"\n"), opt...)
}

// LogInfo logs log line to terminal in yellow
func LogWarning(s string, opt ...interface{}) {
	fmt.Printf(timeWrap("\x1b[33m"+s+"\x1b[0m\n"), opt...)
}

// LogInfo logs log line to terminal in green
func LogSuccess(s string, opt ...interface{}) {
	fmt.Printf(timeWrap("\x1b[32m"+s+"\x1b[0m\n"), opt...)
}

// LogInfo logs log line to terminal in red
func LogError(s string, opt ...interface{}) {
	fmt.Printf(timeWrap("\x1b[31m"+s+"\x1b[0m\n"), opt...)
}

func PrettyExit(code int, s string, opt ...interface{}) {
	LogError(s, opt...)
	os.Exit(code)
}

func timeWrap(s string) string {
	t := time.Now()
	return fmt.Sprintf("[%s] %s", t.Format("2006-01-02 15:04:05"), s)
}
