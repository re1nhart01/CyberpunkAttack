package ws

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type IGatewayHandler interface {
	base.IHandler
	ServiceChannelHandler(context *gin.Context)
}

func GatewayWsRoute(engine *gin.Engine, handler IGatewayHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.GET("/", handler.ServiceChannelHandler)
	}
}
