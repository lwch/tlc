package main

import (
	"github.com/lwch/tlc/cmd/rm"
	"github.com/lwch/tlc/cmd/run"
	"github.com/lwch/tlc/tlcd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tlc",
	Short: "tiny linux container",
	Long:  "tlc is a tiny linux container",
}

func init() {
	run.RegCmd(rootCmd)
	rm.RegCmd(rootCmd)
	tlcd.RegCmd(rootCmd)
}

func main() {
	rootCmd.Execute()
}
