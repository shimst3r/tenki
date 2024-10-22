// Package cmd implements the CLI for tenki.
package cmd

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

import (
	"fmt"
	"os"

	"github.com/shimst3r/tenki/lib"
	"github.com/spf13/cobra"
)

// Language determines the output language.
var Language string

// Location determines the location to query weather information for.
var Location string

// PathToPng determines the path in which to save output as PNG.
var PathToPng string

// UnitOfMeasurement determines the output's unit of measurement
var UnitOfMeasurement lib.Unit

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tenki",
	Short: "A command-line interface for wttr.in/",
	Long: `Tenki is a CLI for querying weather information from wttr.in/ endpoints.

Users can define their location based on various query parameters including
location, nearest point-of-interest, nearest airport, or public domain name.

Supported location types:

    /paris                  # city name
    /~Eiffel+tower          # any location (+ for spaces)
    /Москва                 # Unicode name of any location in any language
    /muc                    # airport code (3 letters)
    /@stackoverflow.com     # domain name
    /94107                  # area codes
    /-78.46,106.79          # GPS coordinates

Supported units of measurement:

    m                       # metric (SI) (used by default everywhere except US)
    u                       # USCS (used by default in US)
    M                       # show wind speed in m/s`,
	Run: GetWeather,
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
	rootCmd.Flags().StringVar(&Language, "language", "en", "Language to generate output in")
	rootCmd.Flags().StringVar(&Location, "location", "", "Query location")
	rootCmd.Flags().StringVar(&PathToPng, "path-to-png", "", "Location to store PNG output")
	rootCmd.Flags().Var(&UnitOfMeasurement, "unit", "Unit of measurement for outputs")
}
