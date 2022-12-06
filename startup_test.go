package startup_test

import (
	"os/exec"
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
func TestStatus(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.False(t, startup.Status())
}

// nolint: paralleltest
func TestUninstall(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.NotNil(t, startup.Uninstall())
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
