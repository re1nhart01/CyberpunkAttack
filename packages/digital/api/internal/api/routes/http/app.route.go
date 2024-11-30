package http

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type IAppHandler interface {
	base.IHandler
	GetHandler(context *gin.Context)
}

func AppRoute(engine *gin.Engine, handler IAppHandler) {
	router := engine.Group(handler.GetPath())
	{
		router.GET("", handler.GetHandler)
	}
}
