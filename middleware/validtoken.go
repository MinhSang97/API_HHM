package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"app/dbutil"
	"app/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}

		if !isValidToken(token) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token không hợp lệ hoặc đã hết hạn"})
			return
		}

		c.Next()
	}
}

func isValidToken(token string) bool {

	db, err := dbutil.ConnectDB()

	var tokens model.Tokens
	err = db.Where("token = ?", token).First(&tokens).Error
	if err != nil {
		return false
	}

	now := time.Now()
	if now.After(tokens.ExpireTime) {
		return false
	}

	return true
}
