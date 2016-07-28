package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
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

func init() {
	repoCmd.AddCommand(listCmd)
}
