package handlers

import (
	ctx "context"
	"fmt"
	"github.com/cyberpunkattack/api/dtos"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/api/repository"
	"github.com/cyberpunkattack/api/wstore"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/pkg/wstorify"
	"net/http"
	"time"

	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

const SESSION_ROUTE = "session"

type ISessionRepo interface {
	CreateNewCustomSessionIM(ctx ctx.Context, args repository.NewSessionIM) (*repository.ActiveSession, error)
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

func (session *SessionHandler) GetMatchUserHistory(context *gin.Context) {
	context.JSON(helpers.GiveTODO())
}

func (session *SessionHandler) GetActiveMatchesHandler(context *gin.Context) {
	context.JSON(helpers.GiveTODO())
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

	newSession, err := session.CreateNewCustomSessionIM(ctx.Background(), createArgs)
	if err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_10, err.Error(), nil))
		return
	}

	sessionMap := map[string]any{
		"session":   newSession,
		"createdAt": time.Now().String(),
	}

	if err := wstore.Store().Global.Store.BroadcastSpecific(
		createArgs.Invites,
		wstore.INVITE_USER_TO_GAME,
		createArgs.CreatorHash,
		wstore.GLOBAL_WS_CHANNEL,
		sessionMap,
		func(event *wstorify.EssentialEvent) error {
			return nil
		}); err != nil {
		context.JSON(helpers.GiveBadRequestCoded(inlineErrors.ERROR_CODE_11, err.Error(), nil))
	} else {
		context.JSON(helpers.GiveOkResponseWithData(sessionMap))
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
