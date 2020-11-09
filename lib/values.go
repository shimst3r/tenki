// Package lib implements common utility functions used in the tenki CLI.
package lib

import "fmt"

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

// Unit defines possible units of measurement passed to the wttr.in API.
type Unit struct {
	symbol string
}

// String is part of the pflag.Value interface.
func (u *Unit) String() string {
	return u.symbol
}

// Set is part of the pflag.Value interface.
func (u *Unit) Set(symbol string) error {
	switch symbol {
	case "m", "u", "M":
		u.symbol = symbol
		return nil
	default:
		return fmt.Errorf("%s is unsupported unit", symbol)
	}
}

// Type is part of the pflag.Value interface.
// TODO: Understand what this method does and implement it accordingly. 🤔
func (u *Unit) Type() string {
	return ""
}
