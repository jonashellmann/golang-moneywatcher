package main

type Configuration struct {
	User string
	Password string
	Database string
}

var configuration Configuration

func ReadConfiguration() (Configuration, error) {
	config := Configuration{User: "", Password: "", Database: ""}

	return config, nil
}
