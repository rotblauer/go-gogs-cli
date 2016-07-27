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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newRepo bool
var newOrgRepo bool
var repoName string
var orgName string

var path string

type repository struct {
	Name string `json:"name"`
}

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage gogs repositories",
	Long:  `Supported operations are .`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Args: ", args)

		if newRepo {
			// create new repo
			if len(args) == 0 { // testes || RotBlauer testes
				// a name for the repo is required
				fmt.Println("Please provide a name for your new repository.")
				return

			}

			if !newOrgRepo {
				// create repo for user
				path = "/user/repos"
				repoName = args[0]
			} else {
				// create repo for organization
				if len(args) < 2 {
					fmt.Println("Usage: [orgname] [reponame]")
					return
				}
				path = "/org/" + args[0] + "/repos"
				repoName = args[1]
			}

			repo := repository{
				Name: repoName,
			}
			client := &http.Client{}

			jsonString, _ := json.Marshal(repo)

			req, _ := http.NewRequest("POST", viper.GetString("api_url")+path, bytes.NewBuffer(jsonString))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", ("token " + viper.GetString("token")))

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(body))
			}

		} else {

			// get all repos owned by authenticated user
			path = "/user/repos"
			client := &http.Client{}
			req, err := http.NewRequest("GET", viper.GetString("api_url")+path, nil)
			req.Header.Set("Authorization", "token "+viper.GetString("token"))
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}

			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(body))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(repoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	repoCmd.Flags().BoolVarP(&newRepo, "new", "n", false, "Create a new repository. A name is required.")
	repoCmd.Flags().BoolVarP(&newOrgRepo, "org", "o", false, "Specify an organization you own.")
}
