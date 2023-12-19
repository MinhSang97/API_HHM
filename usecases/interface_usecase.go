package usecases

import (
	"app/model"
	"context"
)

type LoginUsecase interface {
	Search(ctx context.Context, ApiUser, ApiPassword string) ([]model.Tokens, error)
}

type MeterUsecase interface {
	Search(ctx context.Context, MeterAssetNo, ReceiveTime string) ([]model.DataMeter, error)
}

type MeterTodayUsecase interface {
	Search(ctx context.Context, MeterAssetNo, Start_date, End_date string) ([]model.DataMeter, error)
}
