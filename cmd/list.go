/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"todocli/lib"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List out all of the incomplete tasks on your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := lib.InitDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		var tasks []lib.Task
		if err := db.All(&tasks); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("\nHere are the tasks on your todo list.")
		for _, t := range tasks {
			fmt.Printf("\t%s\n", t)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
