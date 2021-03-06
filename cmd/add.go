/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"todocli/lib"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your Todo List.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := lib.InitDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		t := lib.Task{Title: strings.Join(args, " ")}
		if err := db.Save(&t); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Task added: %s", t)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
