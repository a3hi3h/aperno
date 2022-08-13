package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createExam(ctx *gin.Context) {
	ctx.Request.ParseForm()
	log.Println("Authentication payload")
	log.Println(ctx.Request.Form["examname"])
	log.Println(ctx.Request.Form["examtype"])
	log.Println(ctx.Request.Form["examlevel"])
	log.Println(ctx.Request.Form["examtime"])
	ctx.Redirect(http.StatusFound, "/user")
}

func (server *Server) referExam(ctx *gin.Context) {
	ctx.Request.ParseForm()
	log.Println("Authentication payload")
	log.Println(ctx.Request.Form["examname"])
	log.Println(ctx.Request.Form["examtype"])
	log.Println(ctx.Request.Form["examlevel"])
	log.Println(ctx.Request.Form["examtime"])
	ctx.Redirect(http.StatusFound, "/user")
}
