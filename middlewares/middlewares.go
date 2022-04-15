package middlewares

import (
	token "BookCrud/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {
		err := token.IsTokenValid(context)
		if err != nil {
			context.String(http.StatusUnauthorized, "Unauthorized")
			context.Abort()
			return
		}
		context.Next()
	}
}
