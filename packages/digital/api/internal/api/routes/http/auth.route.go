package http

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	base.IHandler
	SignUpHandler(context *gin.Context)
	ValidatePhoneOrEmailHandler(context *gin.Context)
	LogInHandler(context *gin.Context)
	RefreshTokenHandler(context *gin.Context)
}



func AuthRoute(engine *gin.Engine, handler IAuthHandler) {
	group := engine.Group(handler.GetPath())
	{
		group.POST(SIGN_UP_ROUTE, handler.SignUpHandler)
		group.POST(VALIDATE_ROUTE, handler.ValidatePhoneOrEmailHandler)
		group.POST(LOG_IN_ROUTE, handler.LogInHandler)
		group.POST(REFRESH_ROUTE, handler.RefreshTokenHandler)
	}
}