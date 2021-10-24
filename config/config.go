package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct with plex server info
type Config struct {
	Exporter exporter `json:"exporter"`
	Server server `json:"server"`
}

type exporter struct {
	Port    int    `json:"port"`
}

type server struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Token   string `json:"token"`
}

// Load configuration of plex server from json file and return a Config struct
func Load(file string) (*Config, error) {

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Config file %s does not exist", file)
		}
	}

	data, _ := ioutil.ReadFile(file)

	var config Config
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return nil, fmt.Errorf("Config file %s is not a valid json", file)
	}
	return &config, nil
}
