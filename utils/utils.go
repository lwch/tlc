package utils

import "os"

func ExecDir() string {
	dir, err := os.Executable()
	if err != nil {
		return ""
	}
	return dir
}
