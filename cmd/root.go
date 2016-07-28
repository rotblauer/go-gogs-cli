package cmd

import (
	"fmt"
	"os"

	gogs "github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var apiURL string
var tokenArg string
var client *gogs.Client

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gogs",
	Short: "Connect to the Gogs API.",
	Long: `Welcome to the Gogs CLI.

You'll probably, almost certainly, want a token for interacting with private data.

Visit your Profile Settings on Gogs and *create a token*.

You'll stick that into the default config file named below, 
where you'll also want to *set your base Gogs url*.
	
	$HOME/.go-gogs-cli.yaml 

Recap:
	- get a token from the Gogs UI Profile/settings page.
	- put the token into the config file name above, along with 
	- the base url for your Gogs instance.

There's an example .go-gogs-cli.yaml`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		if token != "" {
			fmt.Println("Token authentication enabled.")
		} else {
			fmt.Println("No token found.")
		}

		url := viper.GetString("api_url")
		if url != "" {
			fmt.Println("Using API url @ ", url)
		} else {
			fmt.Println("No API url coming through... uh oh.")
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initClient)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogs-cli.yaml)")
	RootCmd.PersistentFlags().StringVar(&apiURL, "url", viper.GetString("api_url"), "api url should include /api/v1 path (default is try.gogs.io/api/v1)")
	RootCmd.PersistentFlags().StringVar(&tokenArg, "token", viper.GetString("token"), "token authorization (if not specified in cfg file)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetClient() *gogs.Client {
	// fmt.Println("Getting client....")
	return client
}

func initClient() {
	url := viper.GetString("api_url")
	token := viper.GetString("token")
	client = gogs.NewClient(url, token)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-gogs-cli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")        // adding home directory as first search path
	viper.AutomaticEnv()                // read in environment variables that match
	// viper.SetDefault("api_url", "try.gogs.io/api/v1")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No configuration file found. Is there for sure one at " + viper.ConfigFileUsed() + "?")
	}

}
