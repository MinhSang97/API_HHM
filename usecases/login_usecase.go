package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type loginUseCase struct {
	loginRepo repo.LoginRepo
}

func NewLoginUseCase() LoginUsecase {
	db, _ := dbutil.ConnectDB()
	loginRepo := mysql.NewLginRepository(db)
	return &loginUseCase{
		loginRepo: loginRepo,
	}
}

func (uc *loginUseCase) Search(ctx context.Context, ApiUser, ApiPassword string) ([]model.Tokens, error) {
	return uc.loginRepo.Search(ctx, ApiUser, ApiPassword)
}
