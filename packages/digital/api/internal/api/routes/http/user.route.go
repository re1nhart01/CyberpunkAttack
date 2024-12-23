package http

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	base.IHandler
	GetMyUserProfileHandler(context *gin.Context)
}

func UserRoute(engine *gin.Engine, handler IUserHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.GET(USER_ME_ROUTE, handler.GetMyUserProfileHandler)
	}
}
