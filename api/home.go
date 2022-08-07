package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) homePage(ctx *gin.Context) {
	var tmpName string

	switch ctx.Request.RequestURI {
	case "/":
		tmpName = "home"
	case "/pricing":
		tmpName = "pricing"
	case "/about":
		tmpName = "about"
	case "/login":
		tmpName = "login"
	case "/signup":
		tmpName = "signup"
	}

	ctx.HTML(http.StatusOK, tmpName, gin.H{
		"title": "Aperno Home Page",
	})
}
