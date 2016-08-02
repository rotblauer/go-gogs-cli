package cmd

import (
	"fmt"

	"github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Aliases: []string{"r"},
	Use:     "repo",
	Short:   "parent command for repositories",
	Long: `gogs repo [(new|create)|list|destroy]

	$ gogs repo new my-new-repo --private
	$ gogs repo create my-new-repo --org=JustUsGuys
	$ gogs repo list
	$ gogs repo migrate ia/my-copy-cat https://github.com/gogits/gogs.git
	$ gogs repo destroy ia my-new-repo
	$ gogs repo destroy ia/my-new-repo

	`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Please use: gogs repo [(new|create)|list|destroy]")
	// },
}

func printRepo(repo *gogs.Repository) {
	fmt.Println()
	fmt.Printf("| * %v", repo.FullName)
	if repo.Private {
		fmt.Printf(" (private)")
	}
	if repo.Fork {
		fmt.Printf(" (fork)")
	}
	fmt.Println("")

	if repo.Description != "" {
		fmt.Println("| --> ", repo.Description)
	}

	fmt.Println("| HTML: ", repo.HtmlUrl)
	fmt.Println("| SSH : ", repo.SshUrl)
	fmt.Println("| GIT : ", repo.CloneUrl)
	fmt.Println()
}

func init() {
	RootCmd.AddCommand(repoCmd)
}
