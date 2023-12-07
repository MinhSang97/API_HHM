package handler

import (
	"app/usecases"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Lấy giá trị của các tham số từ query string
		API_User := c.Query("user")
		API_PassWord := c.Query("password")

		fmt.Println("Seach with like", API_User, API_PassWord)

		// Kiểm tra xem có ít nhất một tham số được truyền vào không
		if API_User == "" && API_PassWord == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "At least one search parameter is required",
			})
			return
		}

		uc := usecases.NewLoginUseCase()

		tokens, err := uc.Search(c.Request.Context(), API_User, API_PassWord)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		tokenResponse := tokens[0].Token

		c.JSON(http.StatusOK, gin.H{
			"token": tokenResponse,
		})
	}
}
