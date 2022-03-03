package main

import (
	"os"
	"os/exec"
	"path/filepath"
	rt "runtime"
	"syscall"

	"github.com/lwch/runtime"
	_ "github.com/lwch/tlc/nsenter"
)

func parent() {
	exeDir, err := os.Executable()
	runtime.Assert(err)
	rootfs := filepath.Join(filepath.Dir(exeDir), "..", "debian")
	runtime.Assert(syscall.Chdir(rootfs))
	// TODO: syscall.Exec
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Chroot: rootfs,
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID,
	}
	cmd.Start()
	go func(parent *exec.Cmd) {
		// cmd := exec.Command(exeDir, "mount")
		// cmd.Env = append(os.Environ(),
		// 	fmt.Sprintf("CT_PID=%d", parent.Process.Pid))
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		// runtime.Assert(cmd.Run())
		// cmd := exec.Command("mount")
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		// runtime.Assert(cmd.Run())
		syscall.Mount("proc", filepath.Join(rootfs, "proc"), "proc", 0, "")
	}(cmd)
	defer syscall.Unmount("proc", 0)
	cmd.Wait()
}

func load() {
	// rt.LockOSThread()
	// defer rt.UnlockOSThread()
	runtime.Assert(syscall.Chroot("/home/lwch/tlc/debian"))
	// runtime.Assert(pivotRoot("/home/lwch/tlc/debian"))
	rt.LockOSThread()
	defer rt.UnlockOSThread()
	runtime.Assert(syscall.Mount("proc", "/proc", "proc", 0, ""))
	// pid := os.Getenv("CT_PID")
	// fd, err := syscall.Open(fmt.Sprintf("/proc/%s/ns/%s", pid, "pid"), syscall.O_RDONLY, 0644)
	// runtime.Assert(err)
	// ret, _, msg := syscall.RawSyscall(308, uintptr(fd), 0, 0)
	// if ret != 0 {
	// 	logging.Error("setns: %v", msg)
	// 	return
	// }
}

func main() {
	if len(os.Args) == 1 {
		parent()
		return
	}
	load()
}
