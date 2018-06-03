package main

type Configuration struct {
	db_user string
	db_user_password string
	db_host string
	db_port string
	db_database_name string
}

var configuration Configuration

func ReadConfiguration() (Configuration, error) {
	config := Configuration{db_user: "", db_user_password: "", db_database_name: "" }

	return config, nil
}
