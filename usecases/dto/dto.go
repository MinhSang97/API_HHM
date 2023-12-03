package dto

import (
	"app/payload"
	"time"
)

type Login struct {
	User_Name string
	PassWord  string
}

type Student struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name" validate:"required" gorm:"-"`
	LastName     string    `json:"last_name" validate:"required"`
	Age          int       `json:"age" validate:"required,gt=0"`
	Grade        float32   `json:"grade" validate:"gte=0,lte=10"`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"entrance_date" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Login) ToPayload() *payload.Login {
	login := &payload.Login{
		User_Name: c.User_Name,
		PassWord:  c.PassWord,
	}
	return login
}
