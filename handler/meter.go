package handler

import (
	"app/payload"
	"app/usecases"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MeterHandler struct {
}

func NewMeterHandler() MeterHandler {
	return MeterHandler{}
}

func (l LoginHandler) GetMeter(ginCtx *gin.Context) {
	meterRequest := payload.MetersRequest{}
	// Lấy giá trị của các tham số từ query string

	if err := ginCtx.ShouldBindJSON(&meterRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: fmt.Errorf("Login error: %w", err).Error(),
		})
	}
	fmt.Println("Seach with like", meterRequest.MeterAssetNo, meterRequest.ReceiveTime)

	// Kiểm tra xem có ít nhất một tham số được truyền vào không
	if meterRequest.MeterAssetNo == "" || meterRequest.ReceiveTime == "" {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: errors.New("At least one search parameter is required").Error(),
		})
		return
	}

	uc := usecases.NewLoginUseCase()

	tokens, err := uc.Search(ginCtx.Request.Context(), meterRequest.MeterAssetNo, meterRequest.ReceiveTime)
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
