//go:build !windows

package startup

import (
	"bytes"
	"fmt"
	"os/exec"
)

const _crontab = "crontab"

// Startup set command auto start.
func Startup(command string) error {
	if command == "" {
		return ErrEmptyCommand
	}

	old, err := crontabList()
	if err != nil {
		return err
	}

	command = fmt.Sprintf("@reboot %s\n", command)

	return crontabUpdate(append(old, []byte(command)...))
}

// Include command.
func Include(command string) bool {
	if command == "" {
		return false
	}

	if _, err := exec.LookPath(_crontab); err != nil {
		return false
	}

	data := []byte(command)

	list, err := crontabList()
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

// Remove command.
func Remove(command string) error {
	if command == "" {
		return ErrEmptyCommand
	}

	data := []byte(command)

	old, err := crontabList()
	if err != nil {
		return err
	}

	lines := bytes.Split(old, []byte{'\n'})
	buf := bytes.Buffer{}

	for _, line := range lines {
		if bytes.Contains(line, data) || len(line) == 0 {
			continue
		}

		buf.Write(line)
		buf.Write([]byte{'\n'})
	}

	return crontabUpdate(buf.Bytes())
}

func crontabList() ([]byte, error) {
	if _, err := exec.LookPath(_crontab); err != nil {
		return nil, err
	}

	cmd := exec.Command(_crontab, "-l")

	data, err := cmd.CombinedOutput()

	if len(data) > 0 && data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}

	return data, err
}

func crontabUpdate(data []byte) error {
	cron := exec.Command(_crontab)

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
