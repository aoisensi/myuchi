package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	gConf config
)

type config struct {
	OAuth map[string]struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}
	Hostname string
}

func initConfig() error {
	data, err := ioutil.ReadFile(optCFGName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, &gConf)
	return nil
}
