package cmd

import (
	"fmt"
	"strings"

	"github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
)

// Flags.
var	migrateMirror bool
var migratePrivate bool

func unusableUsage() {
	fmt.Println("Please argue me <user|organization>/<name-of-new-repo> <http://url.of.cloneable.repo.git>")
	return
}
// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Aliases: []string{"m"},
	Use:   "migrate, m",
	Short: "Migrate a repository from a given remote source.",
	Long: `Usage:

$ gogs repo migrate irstacks/copycat https://github.com/gogits/gogs.git
$ gogs repo migrate myCompany/copycat https://github.com/gogits/gogs.git

Options:
[-m | --mirror]  If to be mirrored repo. (no arguments),
[-p | --private] If to be private repo. (no args)
`,
	Run: func(cmd *cobra.Command, args []string) {

		var opts gogs.MigrateRepoOption

		if len(args) != 2 {
			unusableUsage()
			return
		}

		ownerSlashReponame := args[0]
		migrateUrl := args[1]

		// irstacks/my-copy-cat  -->  [irstacks, my-copy-cat]
		ownernameReponame := strings.Split(ownerSlashReponame, "/")
		if len(ownernameReponame) != 2 {
			unusableUsage()
			return
		}

		// get "userId" for owner name from extracurricular function...
		// if the user return err, then we'll need to check if an organization was intended.
		// this is an inconvenience by gogs api client. they should more explicity specify ownership
		// in the migrateRepoOptions.
		user, err := getUserByName(ownernameReponame[0])
		if err != nil {
			// I don't think it actually ever goes here...
			// which means getUserByName works as well for orgs. huh.
			fmt.Println(err)
			fmt.Println("Searching for an org by than name...")

			org, oerr := GetClient().GetOrg(ownernameReponame[0])
			if oerr != nil {
				fmt.Println("Could find neither user nor org by by the name: ", ownernameReponame[0])
				fmt.Println(err)
				return
			} else {
				opts.UID = int(org.ID)
			}
		} else {
			opts.UID = int(user.ID)
		}

		opts.CloneAddr = migrateUrl
	  opts.RepoName = ownernameReponame[1]
		opts.Mirror = migrateMirror
		opts.Private = migratePrivate

		repo, err := GetClient().MigrateRepo(opts)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Repo migrated! Woohoo!")
	  printRepo(repo)
	},
}

func init() {
	repoCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().BoolVarP(&migrateMirror, "mirror", "m", false, " make your migrated repo a mirror of the original")
	migrateCmd.Flags().BoolVarP(&migratePrivate, "private", "p", false, " make your migrated repo private")
}
