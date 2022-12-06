package startup

import (
	"fmt"
	"os"
	"os/exec"
)

// Install startup.
func Install(args ...any) error {
	command, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}

	if Include(command) {
		return nil
	}

	for _, arg := range args {
		command += fmt.Sprintf(" %v", arg)
	}

	return Startup(command)
}

// Status by install.
func Status() bool {
	command, err := exec.LookPath(os.Args[0])
	if err != nil {
		return false
	}

	return Include(command)
}

// Uninstall startup.
func Uninstall() error {
	command, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}

	if !Include(command) {
		return nil
	}

	return Remove(command)
}
