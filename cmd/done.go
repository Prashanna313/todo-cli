/*
Copyright © 2022 Prashanna R
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "marks an item as done",
	Long:    `DoneCmd is used to mark an item as done`,
	Run:     DoneItem,
}

func DoneItem(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		todo.SaveItems(dataFile, items)
	} else {
		log.Println(i, "doesn't match any items")
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
