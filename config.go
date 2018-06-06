package main

import(
	"encoding/json"
	"os"
)

type Configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var configuration Configuration

func ReadConfiguration() (Configuration, error) {
	var config Configuration
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
