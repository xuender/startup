package startup_test

import (
	"fmt"

	"github.com/xuender/startup"
)

// nolint: testableexamples
func Example() {
	fmt.Println(startup.Status())
	fmt.Println(startup.Install("-n", 1))
	fmt.Println(startup.Uninstall())
}
