package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/lwch/runtime"
	_ "github.com/lwch/tlc/nsenter"
)

func parent() {
	exeDir, err := os.Executable()
	runtime.Assert(err)
	rootfs := filepath.Join(filepath.Dir(exeDir), "..", "debian")
	runtime.Assert(syscall.Chdir(rootfs))
	cmd := exec.Command("/proc/self/exe", "child")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID,
	}
	runtime.Assert(cmd.Start())
	defer syscall.Unmount(filepath.Join(rootfs, "proc"), 0)
	runtime.Assert(cmd.Wait())
}

func load() {
	const dir = "/home/lwch/tlc/debian"
	runtime.Assert(os.Chdir(dir))
	runtime.Assert(syscall.Chroot(dir))
	runtime.Assert(syscall.Mount("proc", "/proc", "proc", 0, ""))
	runtime.Assert(syscall.Exec("/bin/bash",
		[]string{"/bin/bash"}, os.Environ()))
}

func main() {
	if len(os.Args) == 1 {
		parent()
		return
	}
	load()
}
