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
	case "/user":
		tmpName = "user"
	case "/setting":
		tmpName = "setting"
	case "/create":
		tmpName = "usercreate"
	case "/exam/create":
		tmpName = "examcreate"
	case "/exam/refer":
		tmpName = "examrefer"
	default:
		tmpName = "404"
	}

	ctx.HTML(http.StatusOK, tmpName, gin.H{
		"title": "Aperno Home Page",
	})
}
