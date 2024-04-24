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

// func runGinServer(config config.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatalf("cannot create server: %v", err) // Format message with error details
// 	}

// 	// err = server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatalf("cannot create server: %v", err)
// 	}
// }
