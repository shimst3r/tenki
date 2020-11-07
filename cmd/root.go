/*
Copyright © 2020 Nils Müller <shimst3r@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/shimst3r/tenki/lib"
	"github.com/spf13/cobra"
)

var cfgFile string
var language string
var location string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tenki",
	Short: "A command-line interface for wttr.in/",
	Long: `Tenki is a CLI for querying weather information from wttr.in/ endpoints.

Users can define their location based on various query parameters including
location, nearest point-of-interest, nearest airport, or public domain name.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cmdGet := &cobra.Command{
		Use:   "get [--location location]",
		Short: "Query weather information",
		Long:  "Query weather information either for your default location or for a specific one.",
		Run: func(cmd *cobra.Command, args []string) {
			lib.GetWeather(language, location, os.Stdout)
		},
	}
	cmdGet.Flags().StringVar(&location, "location", "", "Query location")

	rootCmd.PersistentFlags().StringVar(&language, "language", "en", "Language to generate output in")
	rootCmd.AddCommand(cmdGet)
}
