package tlcd

import (
	"net"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"github.com/lwch/tlc/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// RegCmd register service command
func RegCmd(root *cobra.Command) {
	cfg := &service.Config{
		Name:        "tlcd",
		DisplayName: "tlcd",
		Description: "tiny linux container",
		Arguments:   []string{"service", "run"},
	}

	exeDir, err := os.Executable()
	runtime.Assert(err)
	workDir := filepath.Dir(exeDir)

	svc, err := service.New(&Service{
		WorkDir:    workDir,
		Executable: exeDir,
	}, cfg)
	runtime.Assert(err)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "service control",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "install",
		Short: "register tlc service",
		Run: func(*cobra.Command, []string) {
			runtime.Assert(svc.Install())
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "run tlc service",
		Run: func(*cobra.Command, []string) {
			logging.SetSizeRotate(filepath.Join(workDir, "logs"),
				"tlcd", 10*1024*1024, 7, false)
			runtime.Assert(svc.Run())
		},
	})
	root.AddCommand(cmd)
}

// Service container control service
type Service struct {
	WorkDir    string
	Executable string
	listener   net.Listener
}

func (sv *Service) listen() (net.Listener, error) {
	dir := filepath.Join(sv.WorkDir, "tlcd.sock")
	os.Remove(dir)
	return net.Listen("unix", dir)
}

func (sv *Service) run() {
	l, err := sv.listen()
	runtime.Assert(err)
	defer l.Close()

	sv.listener = l

	svr := grpc.NewServer()
	proto.RegisterServiceServer(svr, sv)
	svr.Serve(l)
}

// Start start service
func (sv *Service) Start(s service.Service) error {
	go sv.run()
	return nil
}

// Stop stop service
func (sv *Service) Stop(s service.Service) error {
	if sv.listener != nil {
		sv.listener.Close()
		os.Remove(filepath.Join(sv.WorkDir, "tlcd.sock"))
	}
	return nil
}
