// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy [username/repo-name]",
	Short: "Destroy a given repo.",
	Long:  `Destroy a repo by username/repo-name. CAREFUL! You won't be asked twice.`,
	Run: func(cmd *cobra.Command, args []string) {

		var owner string
		var repo string

		// ia/testes || ia testes
		if (len(args) == 0) || (len(args) > 2) {
			fmt.Println("Please argue me [username/repo-name] or [username repo-name].")
			return
		}

		// ia/testes
		if len(args) == 1 {
			slasher := strings.Split(args[0], "/")
			if len(slasher) == 2 {
				owner, repo = slasher[0], slasher[1]
			} else {
				fmt.Println("Please argue me [username/repo-name] or [username repo-name].")
				return
			}
		}

		// ia testes
		if len(args) == 2 {
			owner, repo = args[0], args[1]
		}

		if (owner == "") || (repo == "") {
			fmt.Println("Please argue me [username/repo-name] or [username repo-name].")
			return
		}

		err := GetClient().DeleteRepo(owner, repo)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Deleted %v/%v.\n\n", owner, repo)
	},
}

func init() {
	repoCmd.AddCommand(destroyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destroyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destroyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
