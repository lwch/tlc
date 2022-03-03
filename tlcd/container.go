package tlcd

import (
	"os"
	"path/filepath"
	"syscall"

	"github.com/lwch/tlc/proto"
)

type Container struct {
	rootfs      string
	procMounted bool
	devMounted  bool
	sysMounted  bool
	ptsMounted  bool
}

func (ct *Container) pivotRoot(newroot string) error {
	putold := filepath.Join(newroot, "/.pivot_root")

	err := syscall.Mount(newroot, newroot, "", syscall.MS_BIND|syscall.MS_REC, "")
	if err != nil {
		return err
	}

	err = os.MkdirAll(putold, 0700)
	if err != nil {
		return err
	}
	defer os.RemoveAll(putold)

	err = syscall.PivotRoot(newroot, putold)
	if err != nil {
		return err
	}

	err = os.Chdir("/")
	if err != nil {
		return err
	}

	putold = "/.pivot_root"
	return syscall.Unmount(putold, syscall.MNT_DETACH)
}

func (sv *Service) createContainer(cfg *proto.RunConfig) {
}
