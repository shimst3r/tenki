// Package lib implements common utility functions used in the tenki CLI.
package lib

import (
	"io"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://wttr.in/"

// GetWeather writes weather information for location to w.
func GetWeather(language, location string, w io.Writer) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL+location, nil)
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
