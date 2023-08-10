package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanvan229/go23/exercise-06/pkg/model"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		token := strings.TrimSpace(strings.Replace(tokenString, "Bearer", "", 1))

		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		userId, err := model.ValidateToken(token)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set("userId", userId)
		context.Next()
	}
}
