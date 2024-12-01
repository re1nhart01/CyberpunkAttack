package middleware

import "github.com/gin-gonic/gin"

func EncryptionSecurity(context *gin.Context) {

	context.Next()
}