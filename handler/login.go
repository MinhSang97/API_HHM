package handler

import (
	"app/payload"
	"app/usecases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
}

func NewLoginHandler() LoginHandler {
	return LoginHandler{}
}

func (l LoginHandler) Login(ginCtx *gin.Context) {

	loginRequest := payload.UserLoginRequest{}
	// Lấy giá trị của các tham số từ query string

	if err := ginCtx.ShouldBindJSON(&loginRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: fmt.Errorf("Login error: %w", err).Error(),
		})
	}
	fmt.Println("Seach with like", loginRequest.Username, loginRequest.Password)

	// Kiểm tra xem có ít nhất một tham số được truyền vào không
	if loginRequest.Username == "" || loginRequest.Password == "" {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: errors.New("At least one search parameter is required").Error(),
		})
		return
	}

	uc := usecases.NewLoginUseCase()

	tokens, err := uc.Search(ginCtx.Request.Context(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokenResponse := tokens[0].Token

	ginCtx.JSON(http.StatusOK, payload.Response{
		Data: gin.H{
			"token": tokenResponse,
		},
	})

}
