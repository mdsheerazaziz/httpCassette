package server

import (
	"../../utilities"
	"../../config"
	"../../routers"

	"net/http"
	"time"
	"log"
)

// StartServer starts the server
// @todo Start secure server in future
func StartServer() {

	router 		:= routers.InitialiseRoutes()

	srv 		:= &http.Server{
		Addr:           ":" + config.Getenv("SERVER_PORT"),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	utilities.Log("MySQL connection: "+ config.Getenv("CONNECTION_STRING"))
	utilities.Log("Server Starting at port: "+ config.Getenv("SERVER_PORT"))
	log.Fatal(srv.ListenAndServe())
}