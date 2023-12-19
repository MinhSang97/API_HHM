package mysql

import (
	"app/model"
	"app/repo"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type studentRepository struct {
	db *gorm.DB
}

type loginRepository struct {
	db *gorm.DB
}

func (s loginRepository) Search(ctx context.Context, API_User, API_PassWord string) ([]model.Tokens, error) {
	var tokens []model.Tokens

	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15")

	var tokenString string

	tokenString = "hhm@1997" + API_User + API_PassWord + currentDate + currentTime

	hash := sha256.Sum256([]byte(tokenString))
	token := hex.EncodeToString(hash[:])
	fmt.Println(token)

	loginDate := time.Now()
	now := time.Now()

	// Tạo thời gian hết hạn sau 8 tiếng
	expireTime := now.Add(8 * time.Hour)
	expireDate := expireTime
	fmt.Println(loginDate)

	if err := s.db.Where(&model.Tokens{UserName: API_User, Password: API_PassWord}).
		Updates(&model.Tokens{
			Token:      token,
			LoginDate:  loginDate,
			ExpireTime: expireTime,
			ExpireDate: expireDate,
		}).Error; err != nil {
		return nil, fmt.Errorf("insert token error: %w", err)
	}

	// Use ? instead of % in the WHERE clause
	if err := s.db.Where("user_name = ? AND password = ?", API_User, API_PassWord).
		Find(&tokens).Error; err != nil {
		return nil, err
	}

	return tokens, nil
}

var instancelogin loginRepository

func NewLginRepository(db *gorm.DB) repo.LoginRepo {
	if instancelogin.db == nil {
		instancelogin.db = db

	}
	return instancelogin
}
