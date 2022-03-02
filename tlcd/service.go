package tlcd

import (
	"flag"
	"net"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"github.com/lwch/tlc/proto"
	"google.golang.org/grpc"
)

func usage() {
}

// Do run service commands
func Do() {
	set := flag.NewFlagSet("service", flag.PanicOnError)
	help := flag.Bool("h", false, "show help info")
	runtime.Assert(set.Parse(os.Args[2:]))

	if *help {
		usage()
		os.Exit(0)
	}

	cfg := &service.Config{
		Name:        "tlcd",
		DisplayName: "tlcd",
		Description: "tiny linux container",
		Arguments:   []string{"service"},
	}

	exeDir, err := os.Executable()
	runtime.Assert(err)
	workDir := filepath.Dir(exeDir)

	sv, err := service.New(&Service{
		WorkDir: workDir,
	}, cfg)
	runtime.Assert(err)

	switch set.Arg(0) {
	case "install":
		runtime.Assert(sv.Install())
	default:
		logging.SetSizeRotate(filepath.Join(workDir, "logs"), "tlcd", 10*1024*1024, 7, false)
		runtime.Assert(sv.Run())
	}
}

// Service container control service
type Service struct {
	WorkDir  string
	listener net.Listener
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
