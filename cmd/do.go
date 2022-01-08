/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"todocli/lib"

	"github.com/asdine/storm/q"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a task.",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := lib.InitDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("argument provided is not a number, provide a task number")
			return
		}

		var t lib.Task
		query := db.Select(q.Eq("ID", id))
		if err := query.First(&t); err != nil {
			fmt.Println("cannot find task with requested ID", id, err)
			return
		}
		if err := query.Delete(new(lib.Task)); err != nil {
			fmt.Println("problem deleting tasks with requested ID", err)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
