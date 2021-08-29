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
	"strings"
	"task/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long: `Use this to add a task to your TODO list.
Example: task add do the laundry`,
	Run: func(cmd *cobra.Command, args []string) {
		var sb strings.Builder
		for i, arg := range args {
			sb.WriteString(arg)
			if i < len(args)-1 {
				sb.WriteString(" ")
			}
		}
		taskDescription := sb.String()
		store, err := todo.NewStore()
		if err != nil {
			log.Fatal(err)
		}
		defer store.Close()
		err = store.CreateTask(taskDescription)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task added: \"%s\"\n", taskDescription)
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