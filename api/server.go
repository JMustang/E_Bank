package api

import (
	db "github.com/JMustang/E_Bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Define routes
	router.POST("/accounts", server.createAccount)

	server.router = router
	// Return the server instance
	return server
}
