package cmd

import (
	"os"
	"path/filepath"

	"github.com/lwch/runtime"
	"github.com/lwch/tlc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect() (proto.ServiceClient, func(), error) {
	exeDir, err := os.Executable()
	runtime.Assert(err)
	workDir := filepath.Dir(exeDir)
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial("unix://"+filepath.Join(workDir, "tlcd.sock"),
		grpc.WithTransportCredentials(creds))
	runtime.Assert(err)
	return proto.NewServiceClient(conn), func() { conn.Close() }, nil
}
