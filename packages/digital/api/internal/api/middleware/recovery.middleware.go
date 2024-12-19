package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				fmt.Printf("Recovered from panic: %v\n", err)

				// Respond with a custom error message
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error. Please try again later.",
				})
				// Prevent Gin from propagating the panic further
				c.Abort()
			}
		}()
		// Process the request
		c.Next()
	}
}
