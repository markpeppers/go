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
	"strconv"
	"task/todo"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete the task",
	Long:  `Deletes the task fully. Task will no longer be accessible.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskNum, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		store, err := todo.NewStore()
		if err != nil {
			log.Fatal(err)
		}
		defer store.Close()
		desc, err := store.DeleteTask(taskNum)
		fmt.Printf("You have deleted the \"%s\" task.\n", desc)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
