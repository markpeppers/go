/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"task/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show the TODO list",
	Long: `List the numbered tasks in your TODO list. The number
associated with each task can be used to mark "complete" with the
"do" command.`,
	Run: func(cmd *cobra.Command, args []string) {
		store, err := todo.NewStore()
		if err != nil {
			log.Fatal(err)
		}
		defer store.Close()
		tasks, err := store.GetTasks()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%3d. %s\n", i+1, t.Description)
		}
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
