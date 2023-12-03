package handler

import (
	"app/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Lấy giá trị của các tham số từ query string
		// API_UserName := c.Query("user")
		// API_PassWord := c.Query("password")

		API_UserName, _ := c.Request.Context().Value("API_UserName").(string)
		API_PassWord, _ := c.Request.Context().Value("API_Password").(string)

		fmt.Println("Input", API_UserName, API_PassWord)

		// Kiểm tra xem có ít nhất một tham số được truyền vào không
		if API_UserName == "" && API_PassWord == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "At least one search parameter is required",
			})
			return
		}

		// uc := usecases.NewStudentUseCase()

		// students, err := uc.Search(c.Request.Context(), user)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"students": students, // Update the key to match the actual data structure
		// })

	}
}

func InfoLogin(db *gorm.DB) (*model.Login, error) {

	// Lấy username & password

	userinfo := &model.Login{
		user:     API_UserName,
		password: API_PassWord,
	}

	return userinfo, nil

}
