package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	groupName      string
	sessionService SessionService
}

func NewSessionHandler(sessionService SessionService) *SessionHandler {
	return &SessionHandler{"api/session", sessionService}
}

func (handler *SessionHandler) RegisterEndPoints(r *gin.Engine) {
	sessionGroup := r.Group(handler.groupName)

	sessionGroup.POST("", handler.CreateSession)
	sessionGroup.GET("", handler.GetSessions)
	sessionGroup.GET("/:id", handler.GetSession)
	sessionGroup.PUT("/:id", handler.UpdateSession)
	sessionGroup.DELETE("/:id", handler.DeleteSession)
}

func (handler *SessionHandler) CreateSession(ctx *gin.Context) {
	sessionData := NewInputCreateSession()
	err := ctx.BindJSON(&sessionData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed bind session data"})
		return
	}
	newSession, err := handler.sessionService.CreateSession(sessionData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to create new session"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "session created successfully", "data": newSession})
}

func (handler *SessionHandler) GetSessions(ctx *gin.Context) {
	allSessions, err := handler.sessionService.GetSessions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to fetch sessions data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "sessions data fetched successfully", "data": allSessions})
}

func (handler *SessionHandler) GetSession(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid session id"})
		return
	}
	singleSession, err := handler.sessionService.GetSession(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed get session data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "session data fetched scessfully", "data": singleSession})
}

func (handler *SessionHandler) UpdateSession(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid session id"})
		return
	}
	sessionData := NewInputUpdateSession()

	err := ctx.BindJSON(&sessionData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to bind session data"})
		return
	}
	updatedSessionData, err := handler.sessionService.UpdateSession(id, sessionData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to update session data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "session data updated successfully", "data": updatedSessionData})
}

func (handler *SessionHandler) DeleteSession(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid session id"})
		return
	}
	err := handler.sessionService.DeleteSession(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to delete session"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "session deleted successfully"})
}
