package http

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type ISessionHandler interface {
	base.IHandler
}

func SessionRoute(engine *gin.Engine, handler ISessionHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.GET("/", func(context *gin.Context) {

		})
	}
}
