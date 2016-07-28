package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags.
var limit string // max num results
var userName string

func getUserByName(userName string) (user *gogs.User, err error) {
	user, err = GetClient().GetUserInfo(userName)
	if err != nil {
		fmt.Println(err)
		return
	}

	return user, err
}

func searchRepos(uid, query, limit string) ([]*gogs.Repository, error) {

	client := &http.Client{}
	path := "/api/v1/repos/search?q=" + query + "&uid=" + uid + "&limit=" + limit

	repos, err := getParsedResponse(client, "GET", viper.GetString("api_url")+path, nil, nil)

	return repos, err
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Aliases: []string{"s", "find", "f"},
	Use:     "search",
	Short:   "search myquery [-l -u]",
	Long: `Search repos by keyword, with optional flags for [-l | --limit] and [-u | --user] 

$ gogs repo s waldo
$ gogs repo find waldo
$ gogs repo search waldo -l 5
$ gogs repo find waldo --user=johnny --limit=100

Default limit is 10.
Default all users' repos. 

** NOTE that in order to search PRIVATE repos, you'll need to provide the 
USERNAME FLAG

	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Please argue me something to search... like: $ gogs repo search golang")
			fmt.Println("For help, run $ gogs repo search --help")
			return
		}

		// uid := make(chan string)
		var uid string
		var u *gogs.User
		var e error

		// flag for user name, we'll need to get that user's uid
		if userName != "" {
			// go func() {
			u, e = getUserByName(userName)
			if e != nil {
				fmt.Println(e)
				uid = "0"
			}
			uid = strconv.Itoa(int(u.ID))
			// }()
		} else {
			uid = "0"
		}

		repos, err := searchRepos(uid, args[0], limit)
		if err != nil {
			fmt.Println(err)
			return
		}

		var describeUserScope string
		if uid != "0" {
			describeUserScope = u.UserName + "'s repos"
		} else {
			describeUserScope = "all repos"
		}

		fmt.Println("Searching '" + args[0] + "' in " + describeUserScope + "...")
		for _, repo := range repos {
			// get (mostly) empty data in, need to make n more calls to GetRepo... shit.
			splitter := strings.Split(repo.FullName, "/")
			owner := splitter[0]
			reponame := splitter[1]

			r, e := GetClient().GetRepo(owner, reponame)
			if e != nil {
				fmt.Println(e)
				return
			}
			printRepo(r)
		}

	},
}

func init() {
	repoCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&userName, "user", "u", "", "whose repos to search in")
	searchCmd.Flags().StringVarP(&limit, "limit", "l", "10", "limit number of results")
}
