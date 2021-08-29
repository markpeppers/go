/*
Copyright © 2021 Mark Peppers <mpeppers@gmail.com>

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
	"strings"

	"github.com/markpeppers/phone/store"
	"github.com/spf13/cobra"
)

// normalizeCmd represents the normalize command
var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "Normalize numbers store in the database",
	Long: `All numbers in the database will be normalized to the format:
##########`,
	Run: func(cmd *cobra.Command, args []string) {
		store, err := store.NewStore()
		if err != nil {
			log.Fatal(err)
		}
		nums, err := store.GetNumbers()
		if err != nil {
			log.Fatal(err)
		}
		for _, num := range nums {
			err := store.Update(num, normalize(num))
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Numbers normalized!")
	},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// normalizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// normalizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func normalize(num string) string {
	sb := strings.Builder{}
	for _, b := range num {
		_, err := strconv.Atoi(string(b))
		if err != nil {
			continue
		}
		sb.WriteRune(b)
	}
	return sb.String()
}
