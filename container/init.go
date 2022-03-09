package container

import (
	"os"
	"syscall"

	"github.com/lwch/runtime"
)

func createInit() {
	runtime.Assert(os.Chdir(rootfs))
	runtime.Assert(syscall.Chroot(rootfs))
	runtime.Assert(syscall.Mount("proc", "/proc", "proc", 0, ""))
	runtime.Assert(syscall.Exec("/sbin/init",
		[]string{"/sbin/init"}, os.Environ()))
}
