package cmd

import (
	"fmt"
	"os"

	gogs "github.com/gogits/go-gogs-client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var apiURLArg string
var tokenArg string
var client *gogs.Client

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gogs",
	Short: "Connect to the Gogs API.",
	Long: `Welcome to the Gogs CLI.

$ gogs
  --config   path to config file (default $HOME/.go-gogs-cli.yaml; or you can set GOGS_TOKEN and GOGS_URL env vars)
  --config="$HOME/.my-own-gogs-cli-file.yaml"

$ gogs [repo|r]
$ gogs          [create|c|new|n]
                [migrate|m]
                [list]
                [search|find|s|f]
                [delete|destroy|rid|d]
`,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("GOGS_TOKEN")
		if token != "" {
			fmt.Println("Token authentication enabled.")
		} else {
			fmt.Println("No token found.")
		}

		url := viper.GetString("GOGS_URL")
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
	RootCmd.PersistentFlags().StringVar(&apiURLArg, "url", "", "api url should include /api/v1 path (default is try.gogs.io/api/v1)")
	RootCmd.PersistentFlags().StringVar(&tokenArg, "token", "", "token authorization (if not specified in cfg file)")

}

func GetClient() *gogs.Client {
	return client
}

func initClient() {
	url := viper.GetString("GOGS_URL")
	token := viper.GetString("GOGS_TOKEN")
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

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No configuration file found. Is there for sure one at $HOME/.go-gogs-cli.yaml?")
	}

	// These should override any configFile or env vars.
	if tokenArg != "" {
		viper.Set("token", tokenArg)
	}

	if apiURLArg != "" {
		viper.Set("url", apiURLArg)
	}
}
