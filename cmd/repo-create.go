package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
)

// Flags.
var repoName string
var repoDescription string
var repoIsPrivate bool
var orgName string

var repoRemoteName string

var createCmd = &cobra.Command{
	Aliases: []string{"new", "n", "c"},
	Use:     "create",
	Short:   "Create a repository",
	Long: `create [my-new-repo | [-n | --name]] [-d | --desc]  [-org | --org] [-p | --private] [-r | --add-remote]]

	$ gogs repo create my-new-repo
	$ gogs repo new my-new-repo
	$ gogs repo create -n=my-new-repo
	$ gogs repo create JustUsGuys/my-new-repo --desc="a thing with things" -p=true
	$ gogs repo new my-new-repo --private
	$ gogs repo create my-new-repo --add-remote=origin    Will initialize git if not already, too.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Please argue me a name of a repo to create.")
			return
		}
		// Moving to use MyCompany/new-project as way to specify organization instead of
		// --org flag.
		fullRepoName := args
		if len(fullRepoName) == 1 {
			splitter := strings.Split(fullRepoName, "/")
			if len(splitter) == 2 {
				orgName = splitter[0]
				repoName = splitter[1]
			} else {
				repoName = fullRepoName
			}
		} else {
			// we got MyCompany new-project
			orgName = fullRepoName[0]
			repoName = fullRepoName[1] // this'll take 'cactus' of MyCompany cactus spikey --private; spikey will be igored... FIXME?
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

		// add git url as remote to working dir
		if repoRemoteName != "" {
			// get git path
			getGitComm := exec.Command("/usr/bin/which", "git")
			whichGit, err := getGitComm.Output()
			if err != nil {
				fmt.Println(err)
				fmt.Println("...You have git installed, right?")
				return
			}

			whichGitClean := strings.Replace(string(whichGit), "\n", "", 1)

			gitAddRemoteComm := exec.Command(whichGitClean, "remote", "add", repoRemoteName, repo.CloneUrl)
			_, err = gitAddRemoteComm.Output()

			if err != nil {
				// go a step further and let's try init-ing the repo
				//
				var gitInitDone = make(chan bool)

				go func() {
					gitInitComm := exec.Command(whichGitClean, "init")
					gitInit, initErr := gitInitComm.Output()
					if initErr != nil {
						fmt.Println(initErr)
					} else {
						fmt.Println(string(gitInit))
					}
					gitInitDone <- true
				}()

				// wait for gitInitDone
				select {
				case <-gitInitDone:
					// Apparently exec can only call any given command once.
					// https://github.com/golang/go/issues/10305
					gitAddRemoteComm2 := exec.Command(whichGitClean, "remote", "add", repoRemoteName, repo.CloneUrl)
					_, err = gitAddRemoteComm2.Output()
					if err != nil {
						fmt.Println("error adding remote -- ", err.Error())
					} else {
						gitShowRemotesCommand := exec.Command(whichGitClean, "remote", "-v")
						gitShowRemotes, err := gitShowRemotesCommand.Output()
						if err != nil {
							fmt.Println(err.Error())
						}
						fmt.Println(string(gitShowRemotes))
					}
				}
			} else {
				// else there was git already and remote was added
				// fmt.Println(string(addRemote)) // gotcha: adding a remote success returns ""
				// fmt.Println("Remote added as " + repoRemoteName)
				gitShowRemotesCommand2 := exec.Command(whichGitClean, "remote", "-v")
				gitShowRemotes2, err := gitShowRemotesCommand2.Output()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(gitShowRemotes2))
			}
		}

	},
}

func init() {
	repoCmd.AddCommand(createCmd)

	// createCmd.Flags().StringVarP(&repoName, "name", "n", "", "repo name")
	createCmd.Flags().StringVarP(&repoDescription, "desc", "d", "", "repo description")
	createCmd.Flags().BoolVarP(&repoIsPrivate, "private", "p", false, "repo is private")
	// createCmd.Flags().StringVarP(&orgName, "org", "o", "", "organization")

	createCmd.Flags().StringVarP(&repoRemoteName, "add-remote", "r", "", "remote name")
}
