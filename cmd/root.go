package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var recursive bool

var rootCmd = &cobra.Command{
	Use:   "trm",
	Short: "trm allows you to delete a file by moving it to trash so you can recover it later",
	Long: `trm is a command inspired (or meant to "replace" in some cases, rm by moving the files into
	trash so they can be recovered, say bye to deleting files by accident`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("At least 1 argument is required, but 0 where providen")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fileType, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		switch whatType := (fileType.Mode()).IsDir(); whatType {
		case true && recursive:
			fmt.Println("Moving dir as it is  a dir and recursive is specified")
		case false && recursive:
			fmt.Println("User specified recursive but the directory is a file")
		case true && !recursive:
			fmt.Println("File is a dir but no recursivness was specified")
		default:
			fmt.Println("Is a file and not recursive")
		}

	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Allows to delete directories")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
