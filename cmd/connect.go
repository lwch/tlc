package cmd

import (
	"path/filepath"

	"github.com/lwch/runtime"
	"github.com/lwch/tlc/proto"
	"github.com/lwch/tlc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect() (proto.ServiceClient, func(), error) {
	workDir := filepath.Dir(utils.ExecDir())
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial("unix://"+filepath.Join(workDir, "tlcd.sock"),
		grpc.WithTransportCredentials(creds))
	runtime.Assert(err)
	return proto.NewServiceClient(conn), func() { conn.Close() }, nil
}
