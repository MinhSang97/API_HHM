package middleware

import (
	"app/handler"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		API_UserName := handler.Login(&API_UserName) // Set your API_UserName
		API_PassWord := "some_value" // Set your API_Password

		// Add values to the context
		ctx := context.WithValue(c.Request.Context(), "API_UserName", API_UserName)
		ctx = context.WithValue(ctx, "API_Password", API_PassWord)

		// Update the context in the Gin request
		c.Request = c.Request.WithContext(ctx)

		// Call the next handler
		c.Next()
	}
}
