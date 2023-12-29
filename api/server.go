package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

// server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address and port
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse is a helper to format error message in JSON
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
