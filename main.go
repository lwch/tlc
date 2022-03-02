package main

import (
	"os"

	"github.com/lwch/lcmd/lcd"
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
		lcd.Do()
	}
}
