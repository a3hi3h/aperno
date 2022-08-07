package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/a3hi3h/aperno/db/sqlc"
	"github.com/a3hi3h/aperno/token"
	"github.com/a3hi3h/aperno/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) createUserResponse {
	return createUserResponse{
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		ID:        uuid.New(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Hashedpwd: hashedPassword,
		Type:      1,
	}

	user, err := server.sqlStore.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID          `json:"session_id"`
	AccessToken           string             `json:"access_token"`
	AccessTokenExpiresAt  time.Time          `json:"access_token_expires_at"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time          `json:"refresh_token_expires_at"`
	User                  createUserResponse `json:"user"`
}

func (server *Server) loginUserJson(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(req)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.sqlStore.GetUserFromEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Hashedpwd)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.ID.String(),
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.ID.String(),
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.sqlStore.CreateUserSession(ctx, db.CreateUserSessionParams{
		ID:           refreshPayload.ID,
		Userid:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

type postLoginUserRequest struct {
	email    string `form:"email" binding:"required,email"`
	password string `form:"password" binding:"required"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	userReq := &postLoginUserRequest{}
	r := ctx.Request

	if r.Method == http.MethodGet {
		authload, checkAuth := ctx.Get(authorizationPayloadKey)
		if checkAuth {
			authPayload := authload.(*token.Payload)
			log.Println("Authentication payload")
			log.Println(authPayload.Username)
			log.Println(authPayload.ID)

		} else {
			ctx.HTML(http.StatusOK, "login", gin.H{
				"title": "Aperno Home Page",
			})
		}
		return
	}

	if err := ctx.Bind(userReq); err != nil {
		log.Println("Authentication payload binding failure")
		data := map[string]interface{}{
			"StatusCode": http.StatusInternalServerError,
			"Error":      err,
		}
		ctx.HTML(http.StatusInternalServerError, "error", data)
		return
	}

	log.Println("Authentication payload")
	log.Println(userReq.email)

	user, err := server.sqlStore.GetUserFromEmail(ctx, userReq.email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.HTML(http.StatusNotFound, "error", errorResponse(err))
			return
		}
		ctx.HTML(http.StatusInternalServerError, "error", errorResponse(err))
		return
	}

	err = util.CheckPassword(userReq.password, user.Hashedpwd)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "error", errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.ID.String(),
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error", errorResponse(err))
		return
	}

	/*
		refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
			user.ID.String(),
			server.config.RefreshTokenDuration,
		)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "error", errorResponse(err))
			return
		}

		newsession, err := server.sqlStore.CreateUserSession(ctx, db.CreateUserSessionParams{
			ID:           refreshPayload.ID,
			Userid:       user.ID,
			RefreshToken: refreshToken,
			UserAgent:    ctx.Request.UserAgent(),
			ClientIp:     ctx.ClientIP(),
			IsBlocked:    false,
			ExpiresAt:    refreshPayload.ExpiredAt,
		})
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "error", errorResponse(err))
			return
		}
	*/

	ctx.Set(authorizationPayloadKey, accessPayload)
	ctx.Writer.Header().Set(authorizationHeaderKey, authorizationTypeBearer+" "+accessToken)

	/*
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		log.Println(string(body))
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	*/

	log.Println("Auth payload")
	log.Println("Authentication payload")
	log.Println(accessPayload.Username)
	log.Println(accessPayload.ID)

	ctx.Redirect(http.StatusFound, "/login")
}
