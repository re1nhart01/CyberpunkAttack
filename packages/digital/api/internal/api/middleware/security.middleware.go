package middleware

import (
	"github.com/cyberpunkattack/environment"
	"github.com/cyberpunkattack/helpers"
	"github.com/gin-gonic/gin"
	"strings"
)

func EncryptionSecurity(context *gin.Context) {

	context.Next()
}


func ClientBaseSecurity(context *gin.Context) {
	clientSecret := context.GetHeader("X-CLIENT-SECRET")

	if len(clientSecret) < 3 {
		context.AbortWithStatusJSON(helpers.GiveForbidden())
		return
	}

	clientParts := strings.Split(clientSecret," ")

	clientSecretEnv, clientIDEnv := environment.GEnv().GetVariable("CLIENT_SECRET"),
	environment.GEnv().GetVariable("CLIENT_ID")


	if len(clientParts) < 2 && clientIDEnv != clientParts[0] && clientSecretEnv != clientParts[1] {
		context.AbortWithStatusJSON(helpers.GiveForbidden())
		return
	}

	if len(clientParts[1]) < 3 && !strings.HasPrefix(clientParts[1], "@_") {
		context.AbortWithStatusJSON(helpers.GiveForbidden())
		return
	}

	context.Next()
}