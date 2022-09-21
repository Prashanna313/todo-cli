/*
Copyright © 2022 Prashanna R
*/
package cmd

import (
	"log"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `AddCmd is used to create a new todo item to the list
	- set priority using '--priority' or '-p' accepted values are '1, 2, 3'.  For example:
	
	todo-cli add "test" -p 1
	`,
	Run: addItem,
}

func addItem(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}

	er := todo.SaveItems(dataFile, items)
	if er != nil {
		log.Printf("%v", er)
	}

}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Supported Priority: 1, 2, 3")

}
