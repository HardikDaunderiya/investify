package main

import (
	api "investify/api"
	"investify/config"
	"log"
)

func main() {
	config.LoadConfig(".")

	server := api.NewHTTPServer()
	port := config.EnvVars.PORT
	if port == "" {
		port = "5000"
	}
	serverAddr := "127.0.0.1:" + port
	err := server.Start(serverAddr)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}
