package cmd

import (
	"fmt"
	"os"

	"github.com/diegomagdaleno/trm/lib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trm",
	Short: "trm allows you to delete a file by moving it to trash so you can recover it later",
	Long: `trm is a command inspired (or meant to "replace" in some cases, rm by moving the files into
	trash so they can be recovered, say bye to deleting files by accident`,
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		lib.MoveFile(file)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
