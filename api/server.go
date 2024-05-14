package api

import (
	"investify/api/middleware"
	"investify/config"
	"investify/db/adapters"
	db "investify/db/sqlc"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store db.Store
	// store  *db.SQLStore
	router *gin.Engine
}

// runs on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// NewHTTPServer creates a new HTTP server and sets up routing

func NewHTTPServer() *Server {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	pgConn := adapters.InitDb(config.EnvVars.DBSource)

	// if pgConn != nil {
	// 	log.Fatal("Connected to database")
	// }

	store := db.NewStore(pgConn)

	server := &Server{

		router: router,
		store:  store,
	}

	// corsConfig := cors.DefaultConfig()
	log.Println("Setting up server...")

	SetupRouter(server)

	return server
}
