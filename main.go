package main

import (
	"./core/server"
	"./config"
)

// main function
func main() {

	// Initialise .env file
	config.Initialise()

	// Start Server
	server.StartServer()
}