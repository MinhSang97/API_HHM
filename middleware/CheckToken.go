package middleware

import (
	"app/model"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidToken(ctx context.Context, db *gorm.DB, API_User, API_PassWord string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokens []model.Tokens

		if err := db.Where("user_name = ? AND password = ?", API_User, API_PassWord).
			Find(&tokens).
			Error; err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "authentication failed"})
			return
		}

		if len(tokens) > 0 {
			// If authentication is successful, proceed to the next middleware or handler
			// Access the token value (assuming the Tokens model has a field named "Token")
			tokenValue := tokens[0].Token
			fmt.Println("Token value:", tokenValue)

			c.Next()
			return
		}

		c.AbortWithStatusJSON(401, gin.H{"error": "authentication failed, check your username and password"})
	}
}
