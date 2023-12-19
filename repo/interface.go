package repo

import (
	"app/model"
	"context"
)

type LoginRepo interface {
	Search(ctx context.Context, ApiUser, ApiPassword string) ([]model.Tokens, error)
}

type MeterRepo interface {
	Search(ctx context.Context, MeterAssetNo, ReceiveTime string) ([]model.DataMeter, error)
}

type MeterRepoToday interface {
	Search(ctx context.Context, MeterAssetNo, Start_date, End_date string) ([]model.DataMeter, error)
}
