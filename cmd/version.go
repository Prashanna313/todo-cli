package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of todo-cli",
	Long:  `All software has versions. This is todo-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo-cli app v0.0.1")
	},
}
