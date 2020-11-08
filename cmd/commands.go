// Package cmd implements the CLI for tenki.
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

// GetWeather passes command-line arguments to the wttr.in/ wrapper.
func GetWeather(cmd *cobra.Command, args []string) {
	var err error
	writer := os.Stdout

	if PathToPng != "" {
		writer, err = os.Create(PathToPng)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	err = lib.GetWttr(Language, Location, PathToPng, writer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
