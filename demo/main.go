package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	rt "runtime"
	"strings"
	"syscall"

	"github.com/lwch/runtime"
	"golang.org/x/sys/unix"
)

func rootfs() string {
	exeDir, err := os.Executable()
	runtime.Assert(err)
	return filepath.Join(filepath.Dir(exeDir), "..", "debian")
}

func parent() {
	rootfs := rootfs()
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
	fmt.Println("pid=", cmd.Process.Pid)
	runtime.Assert(cmd.Wait())
}

func child() {
	rootfs := rootfs()
	runtime.Assert(os.Chdir(rootfs))
	runtime.Assert(syscall.Chroot(rootfs))
	runtime.Assert(syscall.Mount("proc", "/proc", "proc", 0, ""))
	runtime.Assert(syscall.Exec("/bin/bash",
		[]string{"/bin/bash"}, os.Environ()))
}

func enter(dir, name string) {
	f, err := os.Open(filepath.Join(dir, name))
	runtime.Assert(err)
	defer f.Close()
	runtime.Assert(unix.Setns(int(f.Fd()), 0))
}

func attach(pid string) {
	dir := filepath.Join("/proc", pid, "ns")
	enter(dir, "pid")
	enter(dir, "uts")

	rootfs := rootfs()
	runtime.Assert(syscall.Chroot(rootfs))
	runtime.Assert(os.Chdir("/"))

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runtime.Assert(cmd.Run())
}

func main() {
	rt.LockOSThread()
	defer rt.UnlockOSThread()
	if len(os.Args) == 1 {
		parent()
		return
	}
	switch strings.ToLower(os.Args[1]) {
	case "child":
		child()
	case "attach":
		attach(os.Args[2])
	}
}
