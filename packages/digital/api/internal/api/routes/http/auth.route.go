package http

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	base.IHandler
}



func AuthRoute(engine *gin.Engine, handler IAuthHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.GET("/", func(context *gin.Context) {

		})
	}
}