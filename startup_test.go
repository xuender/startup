package startup_test

import (
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/startup"
)

// nolint: paralleltest
func TestInstall(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.NotNil(t, startup.Install())
}

// nolint: paralleltest
func TestInstall_Has(t *testing.T) {
	assert := assert.New(t)
	patches := gomonkey.ApplyFunc(startup.Include, func(file string) bool {
		return true
	})

	defer patches.Reset()

	assert.Nil(startup.Install())
}

// nolint: paralleltest
func TestInstall_Startup(t *testing.T) {
	assert := assert.New(t)
	patches := gomonkey.ApplyFunc(startup.Startup, func(file string) error {
		return nil
	})

	defer patches.Reset()

	assert.Nil(startup.Install("-d", 1))
}

// nolint: paralleltest
func TestStatus(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.False(t, startup.Status())
}

// nolint: paralleltest
func TestStatus_Include(t *testing.T) {
	patches := gomonkey.ApplyFunc(startup.Include, func(file string) bool {
		return false
	})
	defer patches.Reset()

	assert.False(t, startup.Status())
}

// nolint: paralleltest
func TestUninstall(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.NotNil(t, startup.Uninstall())
}

// nolint: paralleltest
func TestUninstall_Include(t *testing.T) {
	assert := assert.New(t)
	patches := gomonkey.ApplyFunc(startup.Include, func(file string) bool {
		return true
	})
	patches2 := gomonkey.ApplyFunc(startup.Remove, func(file string) error {
		return nil
	})

	defer patches.Reset()
	defer patches2.Reset()

	assert.Nil(startup.Uninstall())
}

// nolint: paralleltest
func TestUninstall_Has(t *testing.T) {
	assert := assert.New(t)
	patches := gomonkey.ApplyFunc(startup.Include, func(file string) bool {
		return false
	})

	defer patches.Reset()

	assert.Nil(startup.Uninstall())
}

func monkeyLookPath() *gomonkey.Patches {
	return gomonkey.ApplyFunc(exec.LookPath, func(file string) (string, error) {
		return "file", startup.ErrEmptyCommand
	})
}

func TestCommandPath(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	path, err := startup.CommandPath()

	assert.Nil(err)
	assert.NotNil(path)
}

// nolint: paralleltest
func TestCommandPath_Abs(t *testing.T) {
	patches := gomonkey.ApplyFunc(filepath.Abs, func(file string) (string, error) {
		return "", startup.ErrEmptyCommand
	})
	defer patches.Reset()

	_, err := startup.CommandPath()

	assert.NotNil(t, err)
}
