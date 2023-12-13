package repo

import (
	"app/model"
	"context"
	"time"
)

type StudentRepo interface {
	GetOneByID(ctx context.Context, id int) (model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
	InsertOne(ctx context.Context, c *model.Student) error
	UpdateOne(ctx context.Context, id int, student *model.Student) error
	DeleteOne(ctx context.Context, id int) error
	Search(ctx context.Context, Value string) ([]model.Student, error)
}

type LoginRepo interface {
	Search(ctx context.Context, API_User, API_PassWord string) ([]model.Tokens, error)
}

type MeterRepo interface {
	Search(ctx context.Context, MeterAssetNo int64, ReceiveTime time.Time) ([]model.Meters, error)
}
