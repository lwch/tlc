package container

import "github.com/spf13/cobra"

var id string
var rootfs string

func RegCmd(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "container",
		Short: "container control",
	}
	cmd.AddCommand(&cobra.Command{
		Use: "init",
		Run: func(cmd *cobra.Command, args []string) {
			createInit()
		},
	})
	cmd.PersistentFlags().StringVar(&id, "id", "", "container id")
	cmd.PersistentFlags().StringVar(&rootfs, "rootfs", "", "rootfs directory")
	root.AddCommand(cmd)
}
