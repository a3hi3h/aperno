package api

import (
	db "github.com/a3hi3h/aperno/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	sqlStore db.Storage
	router   *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(sqlStore db.Storage) *Server {
	server := &Server{sqlStore: sqlStore}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user", server.listUsers)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
