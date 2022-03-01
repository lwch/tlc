package main

import "os"

func usage() {
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
	case "destroy":
	case "start":
	case "stop":
	case "attach":
	}
}
