//go:build windows

package startup

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func Startup(command string) error {
	if command == "" {
		return ErrEmptyCommand
	}

	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	name := GetName(command)

	names, err := key.ReadValueNames(0)
	if err != nil {
		return err
	}

	for _, key := range names {
		if name == key {
			return nil
		}
	}

	return key.SetStringValue(name, GetValue(command))
}

func GetName(command string) string {
	if index := strings.Index(command, " "); index > 0 {
		command = command[:index]
	}

	return filepath.Base(command)
}

func GetValue(command string) string {
	index := strings.Index(command, " ")
	if index < 0 {
		return fmt.Sprintf(`"%s"`, command)
	}

	return fmt.Sprintf(`"%s"%s`, command[:index], command[index:])
}

func Include(command string) bool {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return false
	}
	defer key.Close()

	name := GetName(command)

	names, err := key.ReadValueNames(0)
	if err != nil {
		return false
	}

	for _, key := range names {
		if name == key {
			return true
		}
	}

	return false
}

func Remove(command string) error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.DeleteValue(GetName(command))
}
