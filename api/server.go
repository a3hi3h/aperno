package api

import (
	"fmt"
	"net/http"

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

	router.LoadHTMLGlob("templates/*.tmpl")
	/*
		newtemplates := template.Must(template.ParseFiles(
			"templates/site.tmpl",
			"templates/site.tmpl",
		))
		router.SetHTMLTemplate(newtemplates)
	*/

	router.Static("/assets", "templates/assets")

	router.GET("/", server.homePage)
	router.GET("/about", server.homePage)
	router.GET("/pricing", server.homePage)
	router.GET("/signup", server.homePage)
	router.POST("/login", server.loginUser)

	router.GET("/exam/create", server.homePage)
	router.POST("/exam/create", server.createExam)

	router.GET("/exam/refer", server.homePage)
	router.POST("/exam/refer", server.createExam)

	authCheckRoutes := router.Group("/").Use(authSession(server.tokenMaker, false))
	authCheckRoutes.GET("/login", server.loginUser)
	authCheckRoutes.GET("/user", server.homePage)

	authRoutes := router.Group("/").Use(authSession(server.tokenMaker, true))
	authRoutes.POST("/accounts/create", server.createAccount)
	authRoutes.POST("/accounts", server.getAccount)

	authRoutes.POST("/question/create", server.createQuestion)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404", gin.H{
			"title": "Aperno Home Page",
		})
	})
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
	//router.GET("/*", server.homePage)
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
