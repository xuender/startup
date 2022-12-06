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
	fmt.Println(startup.Has("echo 1"))
	fmt.Println(startup.Startup("echo 1"))
	fmt.Println(startup.Has("echo 1"))
	fmt.Println(startup.End("echo 1"))
	fmt.Println(startup.Has("echo 1"))

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

func TestHas(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.False(startup.Has(""))
	assert.False(startup.Has("test"))
}

// nolint: paralleltest
func TestHas_LookPath(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.False(t, startup.Has("arg"))
}

// nolint: paralleltest
func TestHas_Command(t *testing.T) {
	patches := gomonkey.ApplyFunc(exec.Command, func(name string, args ...string) *exec.Cmd {
		return &exec.Cmd{Stdout: os.Stdout}
	})

	defer patches.Reset()

	assert.False(t, startup.Has("arg"))
}

func TestEnd(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, startup.End(""))
}

// nolint: paralleltest
func TestEnd_LookPath(t *testing.T) {
	defer monkeyLookPath().Reset()

	assert.NotNil(t, startup.End("arg"))
}
