package tlcd

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lwch/runtime"
)

type container struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Rootfs   string `json:"rootfs"`
	Pid      int    `json:"pid"`
	Bootable bool   `json:"bootable"`
	cmd      *exec.Cmd
}

func newContainer(name, rootfs string) *container {
	id, err := runtime.UUID(16, "0123456789abcdef")
	if err != nil {
		return nil
	}
	return &container{
		ID:     id,
		Name:   name,
		Rootfs: rootfs,
	}
}

func (ct *container) save(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	dir = filepath.Join(dir, ct.ID+".json")
	f, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(ct)
}
