package handlers

import (
	"fmt"
	"net/http"

	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

const SESSION_ROUTE = "session"

type ISessionRepo interface {
}

type SessionHandler struct {
	*base.Handler
	ISessionRepo
}

func (session *SessionHandler) GetName() string {
	return session.Name
}

func (session *SessionHandler) GetPath() string {
	return session.Path
}

func (sesion *SessionHandler) CreateSessionHandler(context *gin.Context) {

}

func (user *SessionHandler) GetMyUserProfileHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, map[string]any{
		"status":  http.StatusAccepted,
		"message": "Ok!",
	})
}

func NewSessionHandler(basePath string, repo ISessionRepo) *SessionHandler {
	return &SessionHandler{
		&base.Handler{
			Name: SESSION_ROUTE,
			Path: fmt.Sprintf("/%s/session", basePath),
		},
		repo,
	}
}
