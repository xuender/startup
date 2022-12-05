//go:build !windows && !darwin

package startup_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/startup"
)

func ExampleStartup() {
	fmt.Println(startup.Startup("echo 1"))
	fmt.Println(startup.Has("echo 1"))
	fmt.Println(startup.End("echo 1"))
	fmt.Println(startup.Has("echo 1"))

	// output:
	// <nil>
	// true
	// <nil>
	// false
}

func TestStartup(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, startup.Startup(""))
}

func TestHas(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.False(startup.Has(""))
	ass.False(startup.Has("test"))
}
