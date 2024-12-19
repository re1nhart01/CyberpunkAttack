package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/api/constants"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/api/wstore"
	"github.com/cyberpunkattack/helpers"
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

type IGatewayRepo interface {
}

type GatewayHandler struct {
	*base.Handler
	IGatewayRepo
}

func (gateway *GatewayHandler) GetName() string {
	return gateway.Name
}

func (gateway *GatewayHandler) GetPath() string {
	return gateway.Path
}

func (gateway *GatewayHandler) ServiceChannelHandler(context *gin.Context) {
	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		context.JSON(
			helpers.GiveBadRequestCoded(
				inlineErrors.ERROR_CODE_7,
				"unable to connect to websocket on service",
				nil,
			))
		return
	}
	user := wstorify.NewClient{
		IntoChannel: constants.GLOBAL_CHANNEL,
		User: &wstorify.Account{
			Name:         "test",
			FromGroup:    constants.GLOBAL_CHANNEL,
			Role:         "ADMIN",
			WsConnection: ws,
			UserCredentials: wstore.UserCredentials{
				Id:       1,
				UserHash: "zxczxcz",
				Username: "zxczxc",
			},
			TTC: time.Now(),
		},
	}
	wstore.Store().Global.Register <- &user
	go wstore.ReadGlobalPump(&user, wstore.Store().Global)
	go wstore.ListenGlobal(wstore.Store().Global)
}
func NewGatewayHandler(basePath string, repo IGatewayRepo) *GatewayHandler {
	return &GatewayHandler{
		&base.Handler{
			Name: GATEWAY_ROUTE,
			Path: fmt.Sprintf("/%s/%s", basePath, GATEWAY_ROUTE),
		},
		repo,
	}
}
