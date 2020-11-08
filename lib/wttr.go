// Package lib implements common utility functions used in the tenki CLI.
package lib

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
	"io"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://wttr.in/"

// GetWttr writes weather information from wttr.in/ to io.Writer.
func GetWttr(language, location, pathToPng string, unit Unit, w io.Writer) error {
	client := &http.Client{}
	url := baseURL + location
	if unit.symbol != "" {
		url = url + "?" + unit.symbol
	}
	if pathToPng != "" {
		url = url + ".png"
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept-Language", language)
	req.Header.Add("User-Agent", "curl")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}
