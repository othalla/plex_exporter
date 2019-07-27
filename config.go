package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Token   string `json:"token"`
}

func Load(file string) *Config {
	data, _ := ioutil.ReadFile(file)

	var lol = []byte(data)
	var config Config
	json.Unmarshal(lol, &config)
	return &config
}
