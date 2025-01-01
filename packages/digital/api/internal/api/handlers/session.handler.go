package handlers

import (
	"fmt"
	"github.com/cyberpunkattack/api/dtos"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/api/repository"
	"github.com/cyberpunkattack/helpers"
	"net/http"

	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

const SESSION_ROUTE = "session"

type ISessionRepo interface {
	CreateNewSessionIM(args repository.NewSessionIM) (*repository.ActiveSession, error)
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

func (session *SessionHandler) CreateSessionHandler(context *gin.Context) {
	body, ok := session.Unwrap(context, dtos.CreateSessionDto)
	creds, ok := session.UnwrapUserData(context)
	if !ok {
		return
	}

	createArgs := repository.NewSessionIM{
		Invites:     helpers.AnyToTypeSlice[string](body["invites"].([]any)),
		Password:    helpers.S(body["password"]),
		Name:        helpers.S(body["name"]),
		CreatorHash: helpers.S(creds["userHash"]),
	}

	newSession, err := session.CreateNewSessionIM(createArgs)
	if err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_10, err.Error(), nil))
		return
	}

}

func (session *SessionHandler) GetMyUserProfileHandler(context *gin.Context) {
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
