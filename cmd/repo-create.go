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
	Aliases: []string{"new"},
	Use:     "create",
	Short:   "Create a repository",
	Long: `create [my-new-repo | [-n | --name]] [-d | --desc]  [-org | --org] [-p | --private]]

	$ gogs repo create my-new-repo 
	$ gogs repo new my-new-repo
	$ gogs repo create -n=my-new-repo
	$ gogs repo create my-new-repo --desc="a thing with things" --org=JustUsGuys -p=true
	$ gogs repo new my-new-repo --private`,
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

	createCmd.Flags().StringVarP(&repoName, "name", "n", "", "repo name")
	createCmd.Flags().StringVarP(&repoDescription, "desc", "d", "", "repo description")
	createCmd.Flags().BoolVarP(&repoIsPrivate, "private", "p", false, "repo is private")
	createCmd.Flags().StringVarP(&orgName, "org", "o", "", "organization")
}
