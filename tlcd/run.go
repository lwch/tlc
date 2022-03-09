package tlcd

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/lwch/tlc/proto"
)

// Run handle run command
func (sv *Service) Run(ctx context.Context, cfg *proto.RunConfig) (*proto.RunResponse, error) {
	// TODO: extract image
	ct := newContainer(cfg.GetName(), cfg.Dir)
	if ct == nil {
		return nil, errors.New("can not generate container id")
	}
	sv.Lock()
	sv.cts[ct.ID] = ct
	sv.Unlock()
	cmd := exec.Command(sv.Executable, "container", "init",
		"--id", ct.ID,
		"--rootfs", cfg.Dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID,
	}
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	ct.Pid = cmd.Process.Pid
	ct.cmd = cmd
	err = ct.save(filepath.Join(sv.WorkDir, "containers"))
	if err != nil {
		return nil, err
	}
	go sv.wait(ct)
	return &proto.RunResponse{Id: ct.ID}, nil
}

func (sv *Service) wait(ct *container) {
	ct.cmd.Wait()
	syscall.Unmount(filepath.Join(ct.Rootfs, "proc"), 0)
}
