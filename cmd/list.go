/*
Copyright © 2022 Prashanna R
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items that need to be done",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: ListItems,
}

var (
	allOpt  bool
	doneOpt bool
)

func ListItems(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
		}
	}

	w.Flush()

}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&allOpt, "all", false, "list all items")

}
