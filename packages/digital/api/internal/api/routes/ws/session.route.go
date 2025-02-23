package ws

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type ISessionHandler interface {
	base.IHandler
	GameplayHandler(context *gin.Context)
}

func SessionWsRoute(engine *gin.Engine, handler ISessionHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.GET("/game", handler.GameplayHandler)
	}
}
