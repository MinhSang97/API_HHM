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

type MeterHandlerToday struct {
}

func NewMeterHandlerToday() MeterHandlerToday {
	return MeterHandlerToday{}
}

func (l MeterHandlerToday) GetMeterToday(ginCtx *gin.Context) {
	meterRequestToday := payload.MetersRequestToday{}
	// Lấy giá trị của các tham số từ query string

	if err := ginCtx.ShouldBindQuery(&meterRequestToday); err != nil {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: fmt.Errorf("Get data error: %w", err).Error(),
		})
	}
	fmt.Println("Search with like", meterRequestToday.MeterAssetNo, meterRequestToday.Start_date, meterRequestToday.End_date)

	// Kiểm tra xem có ít nhất một tham số được truyền vào không
	if meterRequestToday.MeterAssetNo == "" || meterRequestToday.Start_date == "" || meterRequestToday.End_date == "" {
		ginCtx.JSON(http.StatusBadRequest, payload.Response{
			Error: errors.New("At least one search parameter is required").Error(),
		})
		return
	}

	//input ngày vào
	inputStart_date := meterRequestToday.Start_date

	t, err := time.Parse("02-01-06", inputStart_date)
	if err != nil {
		panic(err)
	}

	month := t.Month()
	monthStr := month.String()[0:3]

	// Chuyển thành in hoa
	monthStr = strings.ToUpper(monthStr)

	outputStart_date := fmt.Sprintf("%02d-%s-%02d", t.Day(), monthStr, t.Year()%100)

	meterRequestToday.Start_date = outputStart_date

	//output ngày ra
	inputEnd_date := meterRequestToday.Start_date

	t1, err2 := time.Parse("02-01-06", inputEnd_date)
	if err2 != nil {
		panic(err2)
	}

	month1 := t1.Month()
	monthStr1 := month1.String()[0:3]

	// Chuyển thành in hoa
	monthStr = strings.ToUpper(monthStr1)

	outputEnd_date := fmt.Sprintf("%02d-%s-%02d", t.Day(), monthStr, t.Year()%100)

	meterRequestToday.Start_date = outputEnd_date

	uc := usecases.NewMeterTodayUseCase()

	meters, err := uc.Search(ginCtx.Request.Context(), meterRequestToday.MeterAssetNo, meterRequestToday.Start_date, meterRequestToday.End_date)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	meterResponse := meters[0]

	ginCtx.JSON(http.StatusOK, payload.Response{
		Data: gin.H{
			"data": meterResponse,
		},
	})
}
