//go:build windows

package startup_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/startup"
)

func TestGetName(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal("name.exe", startup.GetName("path\\name.exe"))
	assert.Equal("name.exe", startup.GetName("path\\name.exe -d"))
}

func TestGetValue(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal(`"name.exe" -d`, startup.GetValue("name.exe -d"))
	assert.Equal(`"name.exe"`, startup.GetValue("name.exe"))
}
