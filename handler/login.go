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
		user := c.Query("User")
		password := c.Query("Password")

		fmt.Scanln("Input", user, password)

		// Kiểm tra xem có ít nhất một tham số được truyền vào không
		if user == "" && password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "At least one search parameter is required",
			})
			return
		}

		uc := usecases.NewStudentUseCase()

		students, err := uc.Search(c.Request.Context(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"students": students, // Update the key to match the actual data structure
		})
	}
}
