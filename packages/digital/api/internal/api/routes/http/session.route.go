package http

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type ISessionHandler interface {
	base.IHandler
	GetMatchUserHistory(context *gin.Context)
	GetActiveMatchesHandler(context *gin.Context)
	CreateCustomSessionHandler(context *gin.Context)
	ConnectToRandomSession(context *gin.Context)
}

func SessionRoute(engine *gin.Engine, handler ISessionHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.GET("/", handler.GetActiveMatchesHandler)
		group.GET("/:id/history", handler.GetMatchUserHistory)
		group.POST("/new", handler.CreateCustomSessionHandler)
		group.POST("/random", handler.ConnectToRandomSession)
	}
}
