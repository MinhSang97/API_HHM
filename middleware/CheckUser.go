package middleware

import (
	"app/model"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidLogin(ctx context.Context, db *gorm.DB, API_User, API_PassWord string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokens []model.Tokens
		var result int64

		if err := db.Where("user_name = ? AND password = ?", API_User, API_PassWord).
			Find(&tokens).
			Count(&result).Error; err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "authentication failed"})
			return
		}

		if result > 0 {
			// If authentication is successful, proceed to the next middleware or handler
			c.Next()
			return
		}

		c.AbortWithStatusJSON(401, gin.H{"error": "authentication failed, check your username and password"})
	}
}
