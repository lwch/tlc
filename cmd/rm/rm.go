package rm

import "github.com/spf13/cobra"

func RegCmd(root *cobra.Command) {
	root.AddCommand(&cobra.Command{
		Use:   "rm",
		Short: "remove container",
		Run:   do,
	})
}

func do(cmd *cobra.Command, args []string) {
}
