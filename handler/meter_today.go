package handler

import (
	"app/payload"
	"app/usecases"
	"errors"
	"fmt"
	"net/http"
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

	// Format ra định dạng mong muốn
	outputStart_date := t.Format("2006-01-02")

	meterRequestToday.Start_date = outputStart_date

	fmt.Println("inputStart_date", outputStart_date)

	//output ngày ra
	inputEnd_date := meterRequestToday.End_date

	t1, err := time.Parse("02-01-06", inputEnd_date)
	if err != nil {
		panic(err)
	}

	// Format ra định dạng mong muốn
	outputEnd_date := t1.Format("2006-01-02")

	meterRequestToday.Start_date = outputEnd_date

	fmt.Println("outputEnd_date", meterRequestToday.Start_date)

	/////

	uc := usecases.NewMeterTodayUseCase()

	meters, err := uc.Search(ginCtx.Request.Context(), meterRequestToday.MeterAssetNo, meterRequestToday.Start_date, meterRequestToday.End_date)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//meterResponse := meters[0]

	ginCtx.JSON(http.StatusOK, payload.Response{
		Data: gin.H{
			"data": meters,
		},
	})
}
