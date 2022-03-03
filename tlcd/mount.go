package tlcd

import (
	"os"
	"path/filepath"
	"syscall"
)

func (ct *Container) mountProc() error {
	dir := filepath.Join(ct.rootfs, "proc")
	os.RemoveAll(dir)
	err := os.MkdirAll(dir, 0555)
	if err != nil {
		return err
	}
	err = syscall.Mount("proc", dir, "proc",
		syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV, "")
	if err != nil {
		return err
	}
	ct.procMounted = true
	return nil
}

func (ct *Container) mountDev() error {
	dir := filepath.Join(ct.rootfs, "dev")
	os.RemoveAll(dir)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	err = syscall.Mount("tmpfs", dir, "tmpfs",
		syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
	if err != nil {
		return err
	}
	ct.devMounted = true
	return nil
}

func (ct *Container) mountSys() error {
	dir := filepath.Join(ct.rootfs, "sys")
	os.RemoveAll(dir)
	err := os.MkdirAll(dir, 0555)
	if err != nil {
		return err
	}
	err = syscall.Mount("sysfs", dir, "sysfs",
		syscall.MS_NOSUID|syscall.MS_NODEV|syscall.MS_NOEXEC|syscall.MS_RELATIME, "")
	if err != nil {
		return err
	}
	ct.sysMounted = true
	return nil
}

func (ct *Container) mountPts() error {
	dir := filepath.Join(ct.rootfs, "/dev/pts")
	os.RemoveAll(dir)
	err := os.MkdirAll(dir, 0620)
	if err != nil {
		return err
	}
	err = syscall.Mount("devpts", dir, "devpts",
		syscall.MS_NOSUID|syscall.MS_NOEXEC|syscall.MS_RELATIME, "mode=620,ptmxmode=666,gid=5")
	if err != nil {
		return err
	}
	ct.ptsMounted = true
	return nil
}
