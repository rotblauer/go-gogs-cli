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

	"github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
)

// Flags.
var repoName string
var repoDescription string
var repoIsPrivate bool
var orgName string

var createCmd = &cobra.Command{
	Use:   "create [my-new-repo | [-n | --name]] [-d | --desc]  [-org | --org] [-p | --private]]",
	Short: "Create a repository",
	Long:  `Create a repository for you or your organization.`,
	Run: func(cmd *cobra.Command, args []string) {

		// accept gogs repo create thingeys || gogs repo create -name=thingeys
		if repoName == "" {
			repoName = args[0]
		}

		createRepoOpts := gogs.CreateRepoOption{
			Name:        repoName,
			Description: repoDescription,
			Private:     repoIsPrivate,
		}

		var err error
		var repo *gogs.Repository

		if orgName == "" {
			repo, err = GetClient().CreateRepo(createRepoOpts)
		} else {
			repo, err = GetClient().CreateOrgRepo(orgName, createRepoOpts)
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Repo created! Woohoo!")
		printRepo(repo)

	},
}

func init() {
	repoCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&repoName, "name", "n", "", "repo name")
	createCmd.Flags().StringVarP(&repoDescription, "desc", "d", "", "repo description")
	createCmd.Flags().BoolVarP(&repoIsPrivate, "private", "p", false, "repo is private")
	createCmd.Flags().StringVarP(&orgName, "org", "o", "", "organization")
}
