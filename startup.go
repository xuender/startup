package startup

import (
	"fmt"
	"os"
	"os/exec"
)

func Install(args ...any) error {
	command, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}

	if Has(command) {
		return nil
	}

	for _, arg := range args {
		command += fmt.Sprintf(" %v", arg)
	}

	return Startup(command)
}

func Status() bool {
	command, err := exec.LookPath(os.Args[0])
	if err != nil {
		return false
	}

	return Has(command)
}

func Uninstall() error {
	command, err := exec.LookPath(os.Args[0])
	if err != nil {
		return err
	}

	if !Has(command) {
		return nil
	}

	return End(command)
}
