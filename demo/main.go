package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/lwch/runtime"
)

func parent() {
	cmd := exec.Command("/proc/self/exe", "child")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID,
	}
	cmd.Run()
}

func child() {
	runtime.Assert(syscall.Chroot("debian"))
	runtime.Assert(syscall.Chdir("/"))
	runtime.Assert(syscall.Mount("proc", "proc", "proc", 0, ""))
	defer syscall.Unmount("proc", 0)

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID,
	}
	cmd.Run()
}

func main() {
	if len(os.Args) == 1 {
		parent()
		return
	}
	child()
}
