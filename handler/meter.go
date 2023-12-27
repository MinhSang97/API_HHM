package handler

import (
	"app/payload"
	"app/usecases"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MeterHandler struct {
}

func NewMeterHandler() MeterHandler {
	return MeterHandler{}
}

func (l MeterHandler) GetMeter(ginCtx *gin.Context) {
	meterRequest := payload.MetersRequest{}
	// Lấy giá trị của các tham số từ query string

	if err := ginCtx.ShouldBindQuery(&meterRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: fmt.Errorf("Get data error: %w", err).Error(),
		})
	}
	fmt.Println("Search with like", meterRequest.MeterAssetNo, meterRequest.ReceiveTime)

	// Kiểm tra xem có ít nhất một tham số được truyền vào không
	if meterRequest.MeterAssetNo == "" || meterRequest.ReceiveTime == "" {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: errors.New("At least one search parameter is required").Error(),
		})
		return
	}

	str := meterRequest.ReceiveTime

	t, err := time.Parse("02-01-06", str)
	if err != nil {
		panic(err)
	}

	month := t.Month()
	monthStr := month.String()[0:3]

	// Chuyển thành in hoa
	monthStr = strings.ToUpper(monthStr)

	formatted := fmt.Sprintf("%02d-%s-%02d", t.Day(), monthStr, t.Year()%100)

	meterRequest.ReceiveTime = formatted

	fmt.Println("meterRequest", meterRequest)

	uc := usecases.NewMeterUseCase()

	meters, err := uc.Search(ginCtx.Request.Context(), meterRequest.MeterAssetNo, meterRequest.ReceiveTime)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, payload.Response{
		Data: gin.H{
			"data": meters,
		},
	})
}
