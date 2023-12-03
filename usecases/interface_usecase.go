package usecases

import (
	"app/model"
	"context"
)

type StudentUsecase interface {
	GetOneByID(ctx context.Context, id int) (model.Login, error)
	// GetAll(ctx context.Context) ([]model.Student, error)
	// InsertOne(ctx context.Context, c *model.Student) error
	// UpdateOne(ctx context.Context, id int, student *model.Student) error
	// DeleteOne(ctx context.Context, id int) error
	// Search(ctx context.Context, Value string) ([]model.Student, error)
}
