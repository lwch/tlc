package run

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lwch/runtime"
	link "github.com/lwch/tlc/cmd"
	"github.com/lwch/tlc/proto"
	"github.com/spf13/cobra"
)

var name string
var remove bool
var img string
var dir string

// RegCmd register run command
func RegCmd(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a command in a new container",
		Long: `Run a command in a new container, must pass one of --img or --dir argument
  --img argument supported rootfs packed of tar, tar.gz, tar.bz2, zip formats
  --dir argument supported container directory of container.yaml config file`,
		Run: do,
	}
	cmd.Flags().StringVar(&name, "name", "", "Assign a name to the container")
	cmd.Flags().BoolVar(&remove, "rm", false, "Automatically remove the container when it exits")
	cmd.Flags().StringVar(&img, "img", "", "Image file directory")
	cmd.Flags().StringVar(&dir, "dir", "", "Container save directory")
	root.AddCommand(cmd)
}

func do(cmd *cobra.Command, args []string) {
	if len(img) == 0 && len(dir) == 0 {
		fmt.Println("missing --img or --dir")
		os.Exit(1)
	}
	if !filepath.IsAbs(dir) {
		var err error
		dir, err = filepath.Abs(dir)
		runtime.Assert(err)
	}
	cfg := proto.RunConfig{
		Img:  img,
		Dir:  dir,
		Name: name,
		Rm:   remove,
	}
	cli, cancel, err := link.Connect()
	runtime.Assert(err)
	defer cancel()
	resp, err := cli.Run(context.Background(), &cfg)
	runtime.Assert(err)
	fmt.Println(resp.GetId())
}
