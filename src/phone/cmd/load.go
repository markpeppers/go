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
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/markpeppers/phone/store"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load [filename]",
	Short: "Load numbers into database",
	Long: `Load numbers from the file specified.
The file should be a plain text file with one phone
number per line, in any reasonable format.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := "numbers.txt"
		if len(args) > 0 {
			filename = args[0]
		} else {
			log.Println("No file given. Using default \"numbers.txt\"")
		}
		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		store, err := store.NewStore()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(bytes.NewReader(buf))
		writes, errors := 0, 0
		for scanner.Scan() {
			num := scanner.Text()
			if len(num) == 0 {
				continue
			}
			err := store.AddNumber(num)
			if err != nil {
				log.Printf("Error adding number \"%s\": %v", num, err)
				errors++
			} else {
				writes++
			}
		}
		fmt.Printf("Done! %d numbers loaded with %d errors.\n", writes, errors)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
