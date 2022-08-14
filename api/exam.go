package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Exam struct {
	Examname  string
	Examtype  string
	Examlevel string
	Examdesc  string
	Examid    string
	Examtime  int
}

func (server *Server) createExam(ctx *gin.Context) {
	r := ctx.Request
	if r.Method == http.MethodGet {
		ctx.HTML(http.StatusOK, "examrefer", gin.H{
			"title": "Aperno Home Page",
			"examdata": []Exam{
				{Examname: "Software Engineer 2", Examtype: "Python", Examlevel: "Beginer", Examdesc: "Python readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Algo", Examlevel: "Intermediate", Examdesc: "Algo readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Python", Examlevel: "Beginer", Examdesc: "Python readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Algo", Examlevel: "Intermediate", Examdesc: "Algo readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Python", Examlevel: "Beginer", Examdesc: "Python readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Algo", Examlevel: "Intermediate", Examdesc: "Algo readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Python", Examlevel: "Beginer", Examdesc: "Python readability exam", Examid: "uuid", Examtime: 45},
				{Examname: "Software Engineer 2", Examtype: "Algo", Examlevel: "Intermediate", Examdesc: "Algo readability exam", Examid: "uuid", Examtime: 45},
			},
		})
		return
	}
	if r.Method == http.MethodPost {
		ctx.Request.ParseForm()
		log.Println("Authentication payload")
		log.Println(ctx.Request.Form["examname"])
		log.Println(ctx.Request.Form["examtype"])
		log.Println(ctx.Request.Form["examlevel"])
		log.Println(ctx.Request.Form["examtime"])
		ctx.Redirect(http.StatusFound, "/user")
	}
}

func (server *Server) referExam(ctx *gin.Context) {
	ctx.Request.ParseForm()
	log.Println("Authentication payload")
	log.Println(ctx.Request.Form["examname"])
	ctx.Redirect(http.StatusFound, "/user")
}
