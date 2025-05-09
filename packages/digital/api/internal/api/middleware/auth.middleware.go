package middleware

import (
	"strings"

	"github.com/cyberpunkattack/environment"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddlewareHandler(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(helpers.GiveUnauthorized())
		return
	}

	word := environment.GEnv().GetVariable("SPECIAL_WORD_BEARER")
	separator := environment.GEnv().GetVariable("SPECIAL_SYMBOL_BEARER")
	parsedHeader := strings.Split(authHeader, separator)

	if len(parsedHeader) < 2 {
		context.AbortWithStatusJSON(helpers.GiveUnauthorized())
		return
	}
	if parsedHeader[0] != word {
		context.AbortWithStatusJSON(helpers.GiveUnauthorized())
		return
	}

	claims, err := jwt.VerifyToken(strings.TrimSpace(parsedHeader[1]))

	if err != nil {
		context.AbortWithStatusJSON(helpers.GiveUnauthorized())
		return
	}

	body := map[string]any{
		"userHash":       claims.UserHash,
		"id":             claims.Id,
		"expirationTime": claims.ExpiresAt,
		"claims":         claims,
	}

	context.Set("userData", body)
	context.Next()
}
