package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/a3hi3h/aperno/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authSession(tokenMaker token.TokenHandler, required bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader, err := ctx.Request.Cookie(authorizationHeaderKey)

		if err != nil {
			if required {
				ctx.HTML(http.StatusInternalServerError, "error", errorResponse(err))
				return
			}
			ctx.Next()
			return
		}
		log.Println(authorizationHeader.Value)

		/*
			fields := strings.Fields(authorizationHeader.Value)
			if len(fields) < 2 {
				err := errors.New("invalid authorization header format")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
				return
			}

			accessToken := fields[1]
		*/
		payload, err := tokenMaker.VerifyToken(authorizationHeader.Value)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		log.Println(payload.Username)
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.TokenHandler, required bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		log.Println("Auth middleware fpr ")
		if len(authorizationHeader) == 0 {
			if required == false {
				log.Println("authorization header is not provided")
				ctx.Next()
				return
			}
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		log.Println("Setting authorizationPayloadKey ")
		log.Println(payload.Username)
		ctx.Set(authorizationPayloadKey, accessToken)
		ctx.Next()
	}
}
