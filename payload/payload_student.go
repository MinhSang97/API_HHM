package payload

import (
	"app/model"
	"encoding/json"
	"log"
	"time"
)

type AddStudentRequest struct {
	FirstName    string    `json:"first_name" validate:"required"`
	LastName     string    `json:"last_name" validate:"required"`
	Age          int       `json:"age" validate:"required,gt=0"`
	Grade        float32   `json:"grade" validate:"gte=0,lte=10"`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"entrance_date" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AddTokensRequest struct {
	Token      int64     `json:"token"`
	UserName   string    `json:"user_name" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Logindate  time.Time `json:"login_date"`
	ExpireDate time.Time `json:"expire_date"`
	ExpireTime time.Time `json:"expire_time"`
	IPAddress  string    `json:"ip_address"`
}

func (c *AddStudentRequest) ToModel() *model.Student {
	student := &model.Student{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Age:          c.Age,
		Grade:        c.Grade,
		ClassName:    c.ClassName,
		EntranceDate: c.EntranceDate,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}

	return student
}

func (c *AddTokensRequest) ToModel() *model.Tokens {
	token := &model.Tokens{
		Token:      c.Token,
		UserName:   c.UserName,
		Password:   c.Password,
		Logindate:  c.Logindate,
		ExpireDate: c.ExpireDate,
		ExpireTime: c.ExpireTime,
		IPAddress:  c.IPAddress,
	}

	return token
}

func (c *AddStudentRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
