package CIScriptsHelpers

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

// # Timed runs
// def command(*options)
//   log_info(options.join(" "))
//   t = Time.now
//   system(*options)
//   log_success("#{(Time.now - t).round(2)}s\n ")
//   exit $CHILD_STATUS.exitstatus if $CHILD_STATUS && $CHILD_STATUS.exitstatus != 0
// end

// def timed_run(name)
//   log_info(name)
//   t = Time.now
//   yield
//   log_success("#{(Time.now - t).round(2)}s\n ")
// end

// def capture_command(*options)
//   Open3.capture2(*options)
// end

// # system helpers
// def test_command?(*options)
//   system(*options, out: File::NULL)
//   $CHILD_STATUS == 0
// end

// def installed?(binary)
//   test_command?("command", "-v", binary)
// end

type Run func() error

func Command(args ...string) error {
	return TimedRun(strings.Join(args, " "), func() error {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			return err
		}
		err = cmd.Wait()
		return err
	})
}

func TimedRun(name string, function Run) error {
	LogInfo(name)
	start := time.Now()

	err := function()

	end := time.Now()
	runTime := end.Sub(start).Round(time.Millisecond).Seconds()

	logger := LogSuccess
	if err != nil {
		LogError("%s", err)
		logger = LogError
	}

	logger("%fs", runTime)
	return err
}

func CaptureCommand(args ...string) (out []byte, err error) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err = cmd.CombinedOutput()
	return out, err
}

func TestCommand(args ...string) bool {
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	// todo: maybe check child exit status code
	return err == nil
}

func CheckBinary(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
