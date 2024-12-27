package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/api/constants"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/api/wstore"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/pkg/dispatcher"
	"github.com/cyberpunkattack/pkg/wstorify"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const GATEWAY_ROUTE = "gateway"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type IGatewayService interface {
	GetAllSubscriptions() dispatcher.DispatcherSubscription
	UnsubscribeAllEvents() error
}

type IUserRepoInject interface {
	GetUserByField(field string, value any) (*models.UserModel, error)
}

type IGatewayRepo interface {
}

type GatewayHandler struct {
	*base.Handler
	gatewayRepo    IGatewayRepo
	gatewayService IGatewayService
	userRepo       IUserRepoInject
}

func (gateway *GatewayHandler) GetName() string {
	return gateway.Name
}

func (gateway *GatewayHandler) GetPath() string {
	return gateway.Path
}

func (gateway *GatewayHandler) ServiceChannelHandler(context *gin.Context) {
	claims := gateway.UnwrapQueryClaims(context, "token")
	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	close := ws.CloseHandler()
	if err != nil {
		context.JSON(
			helpers.GiveBadRequestCoded(
				inlineErrors.ERROR_CODE_7,
				"unable to connect to websocket on service",
				nil,
			))
		return
	}

	userModel, err := gateway.userRepo.GetUserByField("user_hash", claims.UserHash)
	if err != nil {
		close(helpers.GiveBadWsRequest(websocket.CloseUnsupportedData, "userModel error"))
		return
	}

	user := wstorify.NewClient{
		IntoChannel: constants.GLOBAL_CHANNEL,
		User: &wstorify.Account{
			Name:         userModel.FullName,
			FromGroup:    constants.GLOBAL_CHANNEL,
			Role:         userModel.Role,
			WsConnection: ws,
			UserCredentials: wstore.UserCredentials{
				Id:       userModel.Id,
				UserHash: userModel.UserHash,
				Username: userModel.Username,
			},
			TTC: time.Now(),
		},
	}

	wstore.Store().Global.Register <- &user
	go wstore.ReadGlobalPump(&user, wstore.Store().Global)
}

type GatewayHandlerArgs struct {
	BasePath string
	Repo     IGatewayRepo
	UserRepo IUserRepoInject
	Services IGatewayService
}

func NewGatewayHandler(args GatewayHandlerArgs) *GatewayHandler {
	return &GatewayHandler{
		&base.Handler{
			Name: GATEWAY_ROUTE,
			Path: fmt.Sprintf("/%s/%s", args.BasePath, GATEWAY_ROUTE),
		},
		args.Repo,
		args.Services,
		args.UserRepo,
	}
}
