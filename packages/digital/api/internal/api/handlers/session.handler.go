package handlers

import (
	ctx "context"
	"fmt"
	"github.com/cyberpunkattack/api/constants"
	"github.com/gorilla/websocket"
	"net/http"
	"time"

	"github.com/cyberpunkattack/api/dtos"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/api/repository"
	"github.com/cyberpunkattack/api/wstore"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/pkg/wstorify"

	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

const SESSION_ROUTE = "session"

type ISessionRepo interface {
	CreateNewCustomSessionIM(ctx ctx.Context, args repository.NewSessionIM) (*repository.ActiveSession, error)
}

type ISessionUserRepoInject interface {
	GetUserByField(field string, value any) (*models.UserModel, error)
}

type SessionHandler struct {
	*base.Handler
	ISessionRepo
	userRepo ISessionUserRepoInject
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

func (session *SessionHandler) ConnectToRandomSession(context *gin.Context) {
	context.JSON(helpers.GiveTODO())
}

func (session *SessionHandler) StartGameHandler(context *gin.Context) {
	
}

func (session *SessionHandler) CreateCustomSessionHandler(context *gin.Context) {
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
		Type:        models.GAME_TYPE_CUSTOM,
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

func (session *SessionHandler) GameplayHandler(context *gin.Context) {
	claims := session.UnwrapQueryClaims(context, "token")
	sessionId := context.Query("session_id")
	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	close := ws.CloseHandler()

	if err != nil {
		context.JSON(
			helpers.GiveBadRequestCoded(
				inlineErrors.ERROR_CODE_13,
				"unable to connect to websocket session on service",
				nil,
			))
		return
	}

	userModel, err := session.userRepo.GetUserByField("user_hash", claims.UserHash)
	if err != nil {
		close(helpers.GiveBadWsRequest(websocket.CloseUnsupportedData, "userModel error"))
		return
	}

	user := wstorify.NewClient{
		IntoChannel: constants.SESSION_CHANNEL,
		User: &wstorify.Account{
			Name:         userModel.FullName,
			FromGroup:    constants.SESSION_CHANNEL,
			Role:         userModel.Role,
			WsConnection: ws,
			UserCredentials: wstore.UserCredentials{
				Id:       userModel.Id,
				UserHash: userModel.UserHash,
				Username: userModel.Username,
			},
			TTC: time.Now(),
		},
		Payload: map[string]any{
			"session_id": sessionId,
		},
	}

	wstore.Store().Sessions.Register <- &user
	go wstore.ReadSessionPump(&user, wstore.Store().Sessions)
}

type SessionHandlerArgs struct {
	BasePath string
	Repo     ISessionRepo
	UserRepo ISessionUserRepoInject
}

func NewSessionHandler(args SessionHandlerArgs) *SessionHandler {
	return &SessionHandler{
		&base.Handler{
			Name: SESSION_ROUTE,
			Path: fmt.Sprintf("/%s/session", args.BasePath),
		},
		args.Repo,
		args.UserRepo,
	}
}
