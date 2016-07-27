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

var repoRootCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage your repositories.",
	Long:  "Definitely manage your repositories.",
	Run: func(cmd *cobra.Command, args []string) {
		// do nothing?

		// debug...
		client := GetClient()
		fmt.Printf("\nClient: %v\n", client)

		repos, err := client.ListMyRepos()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repo := range repos {
			printRepo(repo)
		}
	},
}

var repoListCmnd = &cobra.Command{
	Use:   "list",
	Short: "List your repositories.",
	Long:  "List all your repositories.",
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := GetClient().ListMyRepos()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, repo := range repos {
			printRepo(repo)
		}
	},
}

// repoCmd represents the repo command
var repoCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a repository",
	Long:  `Create a repository for you or your organization.`,
	Run: func(cmd *cobra.Command, args []string) {

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

func printRepo(repo *gogs.Repository) {
	fmt.Printf("Name: %v\nGit url: %v\n\n", repo.FullName, repo.CloneUrl)
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	repoCreateCmd.Flags().StringVarP(&repoName, "name", "n", "", "repo name")
	repoCreateCmd.Flags().StringVarP(&repoDescription, "desc", "d", "", "repo description")
	repoCreateCmd.Flags().BoolVarP(&repoIsPrivate, "private", "p", false, "repo is private")
	repoCreateCmd.Flags().StringVarP(&orgName, "org", "o", "", "organization")

	RootCmd.AddCommand(repoRootCmd)
	repoRootCmd.AddCommand(repoCreateCmd)
}

// fmt.Println("Args: ", args)

// if createUserRepoFlag {
// 	// create new repo
// 	if len(args) == 0 { // testes || RotBlauer testes
// 		// a name for the repo is required
// 		fmt.Println("Please provide a name for your new repository.")
// 		return

// 	}

// 	if !createOrganizationRepoFlag {
// 		// create repo for user
// 		path = "/user/repos"
// 		repoName = args[0]
// 	} else {
// 		// create repo for organization
// 		if len(args) < 2 {
// 			fmt.Println("Usage: [orgname] [reponame]")
// 			return
// 		}
// 		path = "/org/" + args[0] + "/repos"
// 		repoName = args[1]
// 	}

// 	repo := repository{
// 		Name: repoName,
// 	}
// 	client := &http.Client{}

// 	jsonString, _ := json.Marshal(repo)

// 	req, _ := http.NewRequest("POST", viper.GetString("api_url")+path, bytes.NewBuffer(jsonString))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", ("token " + viper.GetString("token")))

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(body))
// 	}

// } else {

// 	// get all repos owned by authenticated user
// 	path = "/user/repos"
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", viper.GetString("api_url")+path, nil)
// 	req.Header.Set("Authorization", "token "+viper.GetString("token"))
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(body))
// 	}
// }
