package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type meterTodayUseCase struct {
	meterTodayRepo repo.MeterRepoToday
}

func NewMeterTodayUseCase() MeterTodayUsecase {
	db, _ := dbutil.ConnectDB()
	meterTodayRepo := mysql.NewMeterTodayRepository(db)
	return &meterTodayUseCase{
		meterTodayRepo: meterTodayRepo,
	}
}

func (uc *meterTodayUseCase) Search(ctx context.Context, MeterAssetNo, Start_date, End_date string) ([]model.DataMeter, error) {
	return uc.meterTodayRepo.Search(ctx, MeterAssetNo, Start_date, End_date)
}
