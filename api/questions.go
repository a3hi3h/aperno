package api

import (
	"encoding/json"
	"net/http"

	db "github.com/a3hi3h/aperno/db/sqlc"
	"github.com/a3hi3h/aperno/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type QuestionFormat struct {
	Name   string `json:"name" binding:"required,max=255"`
	Detail string `json:"type" binding:"required,oneof=post nested"`
}
type createQuestionsRequest struct {
	Type     int32           `json:"type" binding:"required"`
	Level    int32           `json:"level" binding:"required"`
	Question json.RawMessage `json:"question" binding:"required"`
}

func (server *Server) createQuestion(ctx *gin.Context) {
	var req createQuestionsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if !json.Valid(req.Question) {
		ctx.JSON(http.StatusBadRequest, req.Question)
		return
	}

	//jsons, err := json.Marshal(req.Question)

	// Check if user exists
	arg := db.CreateQuestionsParams{
		ID:     uuid.New(),
		Type:   req.Type,
		Level:  req.Level,
		Userid: uuid.MustParse(authPayload.Username),
	}

	question, err := server.sqlStore.CreateQuestions(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}
