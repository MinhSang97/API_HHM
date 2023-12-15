package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type meterUseCase struct {
	meterRepo repo.MeterRepo
}

func NewMeterUseCase() MeterUsecase {
	db, _ := dbutil.ConnectDB()
	meterRepo := mysql.NewMeterRepository(db)
	return &meterUseCase{
		meterRepo: meterRepo,
	}
}

func (uc *meterUseCase) Search(ctx context.Context, MeterAssetNo, ReceiveTime string) ([]model.Meters, error) {
	return uc.meterRepo.Search(ctx, MeterAssetNo, ReceiveTime)
}
