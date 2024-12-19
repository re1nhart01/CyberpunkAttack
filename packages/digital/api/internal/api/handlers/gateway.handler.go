package handlers

import (
	"fmt"
	"time"

	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/api/constants"
	"github.com/cyberpunkattack/api/criegstore"
	inlineErrors "github.com/cyberpunkattack/api/errors"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/pkg/crieg"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const GATEWAY_ROUTE = "gateway"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
	defer ws.Close()

	user := crieg.CriegNewClient{
		IntoChannel: constants.GLOBAL_CHANNEL,
		User: &crieg.CriegUser{
			Name:         "test",
			FromGroup:    constants.GLOBAL_CHANNEL,
			Role:         "ADMIN",
			WsConnection: ws,
			UserCredentials: criegstore.CriegUserCredentials{
				Id:       1,
				UserHash: "zxczxcz",
				Username: "zxczxc",
			},
			TTC: time.Now(),
		},
	}
	criegstore.Crieg().Global.Register <- &user

}
func NewGatewayHandler(basePath string, repo IUserRepo) *GatewayHandler {
	return &GatewayHandler{
		&base.Handler{
			Name: GATEWAY_ROUTE,
			Path: fmt.Sprintf("/%s/%s", basePath, GATEWAY_ROUTE),
		},
		repo,
	}
}
