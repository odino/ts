package main

import (
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf(" \n")
	fmt.Printf("  %s         # get current timestamp\n", os.Args[0])
	fmt.Printf("  %s +2h     # get timestamp 2h in the future\n", os.Args[0])
	fmt.Printf("  %s -30s    # get timestamp 30s in the past\n", os.Args[0])
	fmt.Printf(" \n")
	fmt.Printf(" Supported formats https://pkg.go.dev/time#ParseDuration\n")
}

// The dumbest utility you can think of
// to generate unix timestamps on-the-fly.
func main() {
	interval := "0s"

	// if args are specified, take the first one
	// as the interval to generate the timestamp with
	if len(os.Args) > 1 {
		if _, ok := map[string]bool{"-h": true, "--help": true}[os.Args[1]]; ok {
			usage()
			return
		}

		interval = os.Args[1]
	}

	// parse the interval
	d, err := time.ParseDuration(interval)

	// I got 99 problems...
	if err != nil {
		fmt.Println(err.Error())
		usage()
		os.Exit(99)
	}

	fmt.Println(time.Now().Add(d).Unix())

}
