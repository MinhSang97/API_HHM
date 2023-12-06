package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/mysql"
	"context"
)

type studentUseCase struct {
	studentRepo repo.StudentRepo
}

func NewStudentUseCase() StudentUsecase {
	db, _ := dbutil.ConnectDB()
	studentRepo := mysql.NewStudentRepository(db)
	return &studentUseCase{
		studentRepo: studentRepo,
	}
}

func (uc *studentUseCase) GetOneByID(ctx context.Context, id int) (model.Student, error) {

	return uc.studentRepo.GetOneByID(ctx, id)
}

func (uc *studentUseCase) GetAll(ctx context.Context) ([]model.Student, error) {
	return uc.studentRepo.GetAll(ctx)
}

func (uc *studentUseCase) InsertOne(ctx context.Context, student *model.Student) error {
	return uc.studentRepo.InsertOne(ctx, student)
}

func (uc *studentUseCase) UpdateOne(ctx context.Context, id int, student *model.Student) error {
	return uc.studentRepo.UpdateOne(ctx, id, student)
}

func (uc *studentUseCase) DeleteOne(ctx context.Context, id int) error {
	return uc.studentRepo.DeleteOne(ctx, id)
}

func (uc *studentUseCase) Search(ctx context.Context, Value string) ([]model.Student, error) {
	return uc.studentRepo.Search(ctx, Value)
}

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

func (uc *loginUseCase) Search(ctx context.Context, API_User, API_PassWord string) ([]model.Tokens, error) {
	return uc.loginRepo.Search(ctx, API_User, API_PassWord)
}
