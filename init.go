package main

import "go-gin-template/src/config"

func init() {
	// Get the env profile
	config.GetEnvironmentConfig()
	// Connect to DB
	config.CreateDBConnection()
}
