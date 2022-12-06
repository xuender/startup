package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/xuender/startup"
)

func main() {
	install := flag.Bool("i", false, "Install flag example.")
	uninstall := flag.Bool("u", false, "Uninstall flag example.")
	status := flag.Bool("s", false, "Flag example status.")
	daemon := flag.Bool("d", false, "Flag example daemon.")

	flag.Parse()

	switch {
	case *install:
		if err := startup.Install("-d"); err != nil {
			panic(err)
		}

		fmt.Fprintln(os.Stdout, "Flag example install is ok.")
	case *uninstall:
		if err := startup.Uninstall(); err != nil {
			panic(err)
		}

		fmt.Fprintln(os.Stdout, "Flag example uninstall is ok.")
	case *status:
		if startup.Status() {
			fmt.Fprintln(os.Stdout, "Flag example is install.")
		} else {
			fmt.Fprintln(os.Stdout, "Flag example is not install.")
		}
	case *daemon:
		fmt.Fprintln(os.Stdout, "Flag example daemon...")
		time.Sleep(time.Hour)
	default:
		flag.Usage()
	}
}
