package main

import (
	"os"

	"github.com/lwch/tlc/tlcd"
)

func usage() {
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "create":
	case "destroy":
	case "start":
	case "stop":
	case "attach":
	case "service":
		tlcd.Do()
	}
}
