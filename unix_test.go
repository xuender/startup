//go:build !windows && !darwin

package startup_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/startup"
)

func ExampleStartup() {
	fmt.Println(startup.Include("echo 1"))
	fmt.Println(startup.Startup("echo 1"))
	fmt.Println(startup.Include("echo 1"))
	fmt.Println(startup.Remove("echo 1"))
	fmt.Println(startup.Include("echo 1"))

	// Output:
	// false
	// <nil>
	// true
	// <nil>
	// false
}

func TestStartup(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, startup.Startup(""))
}

// nolint: paralleltest
func TestStartup_LookPath(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.NotNil(t, startup.Startup("arg"))
}

func TestInclude(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.False(startup.Include(""))
	assert.False(startup.Include("test"))
}

// nolint: paralleltest
func TestInclude_LookPath(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.False(t, startup.Include("arg"))
}

// nolint: paralleltest
func TestInclude_Command(t *testing.T) {
	patches := gomonkey.ApplyFunc(exec.Command, func(name string, args ...string) *exec.Cmd {
		return &exec.Cmd{Stdout: os.Stdout}
	})

	defer patches.Reset()

	assert.False(t, startup.Include("arg"))
}

func TestRemove(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, startup.Remove(""))
}

// nolint: paralleltest
func TestRemove_LookPath(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.NotNil(t, startup.Remove("arg"))
}
