//go:build !windows && !darwin

package startup

import (
	"bytes"
	"fmt"
	"os/exec"

	"golang.org/x/exp/slices"
)

func Startup(command string) error {
	if command == "" {
		return ErrEmptyCommand
	}

	old, err := list()
	if err != nil {
		return err
	}

	command = fmt.Sprintf("@reboot %s\n", command)

	return update(append(old, []byte(command)...))
}

func Has(command string) bool {
	if command == "" {
		return false
	}

	if _, err := exec.LookPath("crontab"); err != nil {
		return false
	}

	data := []byte(command)

	list, err := list()
	if err != nil {
		return false
	}

	for _, line := range bytes.Split(list, []byte{'\n'}) {
		if bytes.Contains(line, data) {
			return true
		}
	}

	return false
}

func End(command string) error {
	if command == "" {
		return ErrEmptyCommand
	}

	data := []byte(command)

	old, err := list()
	if err != nil {
		return err
	}

	lines := bytes.Split(old, []byte{'\n'})

	for index, line := range lines {
		if bytes.HasSuffix(line, data) {
			lines = slices.Delete(lines, index, index+1)

			break
		}
	}

	return update(bytes.Join(lines, []byte{'\n'}))
}

func list() ([]byte, error) {
	if _, err := exec.LookPath("crontab"); err != nil {
		return nil, err
	}

	cmd := exec.Command("crontab", "-l")

	data, err := cmd.CombinedOutput()

	if len(data) > 0 && data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}

	return data, err
}

func update(data []byte) error {
	cron := exec.Command("crontab")

	pipe, err := cron.StdinPipe()
	if err != nil {
		return err
	}

	if _, err := pipe.Write(data); err != nil {
		return err
	}

	if err := cron.Start(); err != nil {
		return err
	}

	if err := pipe.Close(); err != nil {
		return err
	}

	return cron.Wait()
}
