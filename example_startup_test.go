package startup_test

import (
	"fmt"

	"github.com/xuender/startup"
)

func Example() {
	fmt.Println(startup.Status())
	fmt.Println(startup.Install("-n", 1))
	fmt.Println(startup.Status())
	fmt.Println(startup.Uninstall())
	fmt.Println(startup.Status())

	// Output:
	// false
	// <nil>
	// true
	// <nil>
	// false
}
