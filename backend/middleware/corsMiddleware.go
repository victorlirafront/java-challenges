package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusOK)
			return
		}

		context.Next()
	}
}
