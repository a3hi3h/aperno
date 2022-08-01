package api

import (
	"fmt"

	db "github.com/a3hi3h/aperno/db/sqlc"
	"github.com/a3hi3h/aperno/token"
	"github.com/a3hi3h/aperno/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	sqlStore   db.Storage
	router     *gin.Engine
	tokenMaker token.TokenHandler
	config     util.Config
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, sqlStore db.Storage) (*Server, error) {
	tokenMaker, err := token.NewPasetoToken(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token handler: %w", err)
	}

	server := &Server{
		config:     config,
		sqlStore:   sqlStore,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users", server.listUsers)

	/*
		router.POST("/users/login", server.loginUser)

		authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

			router.POST("/users/login", server.loginUser)
			router.POST("/tokens/renew_access", server.renewAccessToken)

			authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
			authRoutes.POST("/accounts", server.createAccount)
			authRoutes.GET("/accounts/:id", server.getAccount)
			authRoutes.GET("/accounts", server.listAccounts)

			authRoutes.POST("/transfers", server.createTransfer)
	*/
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
