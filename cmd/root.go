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

You'll probably want a token for interacting with private data. 
Visit your profile settings on Gogs and create a token, then stick it into
$HOME/.gogs-cli.yaml. While you're there you can adjust the default to your liking, too.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		if token != "" {
			fmt.Println("Token authentication enabled @ ", token)
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

func GetClient() *gogs.Client {
	fmt.Println("Getting client.")
	return client
}

func initClient() {
	fmt.Println("initClient")

	url := viper.GetString("api_url")
	token := viper.GetString("token")
	fmt.Printf("api url: %v\n", url)
	fmt.Printf("token: %v\n", token)

	client = gogs.NewClient(url, token)
}

func init() {
	cobra.OnInitialize(initConfig, initClient)

	// url := viper.GetString("api_url")
	// toke := viper.GetString("token")
	// fmt.Printf("api url: %v", url)
	// fmt.Printf("token: %v", toke)

	// client = gogs.NewClient(url, toke)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogs-cli.yaml)")
	// RootCmd.PersistentFlags().StringVar(&apiURL, "url", "", "api url should include /api/v1 path (default is try.gogs.io/api/v1)")
	// RootCmd.PersistentFlags().StringVar(&tokenArg, "token", "", "token authorization (if not specified in cfg file)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gogs-cli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match
	// viper.SetDefault("api_url", "try.gogs.io/api/v1")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No configuration file found.")
	}

	// if api_url was flagged, set it to that (override cfg)
	// if apiURL != "" {
	// 	viper.Set("api_url", apiURL)
	// }

	// // if token was flagged, set it to that (override cfg)
	// if tokenArg != "" {
	// 	viper.Set("token", tokenArg)
	// }

}
