package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Aliases: []string{"d", "delete", "rid"},
	Use:     "destroy [username repo-name | username/repo-name]",
	Short:   "Destroy a repo.",
	Long: `destroy [username repo-name | username/repo-name]

	$ destroy ia tester-repo
	$ destroy ia/tester-repo

	**CAREFUL!** YOU WON'T BE ASKED TWICE.
	YE BE WARNED.

	`,
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
}
