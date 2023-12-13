package dto

import (
	"app/payload"
	"time"
)

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

func (c *Student) ToPayload() *payload.AddStudentRequest {
	studentPayload := &payload.AddStudentRequest{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Age:          c.Age,
		Grade:        c.Grade,
		ClassName:    c.ClassName,
		EntranceDate: c.EntranceDate,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
	return studentPayload
}

type Tokens struct {
	Token      string    `json:"token"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	Logindate  time.Time `json:"login_date"`
	ExpireDate time.Time `json:"expire_date"`
	ExpireTime time.Time `json:"expire_time"`
	IPAddress  string    `json:"ip_address"`
}

func (c *Tokens) ToPayload() *payload.AddTokensRequest {
	tokenPayload := &payload.AddTokensRequest{
		Token:      c.Token,
		UserName:   c.UserName,
		Password:   c.Password,
		Logindate:  c.Logindate,
		ExpireDate: c.ExpireDate,
		ExpireTime: c.ExpireTime,
		IPAddress:  c.IPAddress,
	}
	return tokenPayload
}

type Meters struct {
	MeterAssetNo int64     `json:"meter_asset_no"`
	ReceiveTime  time.Time `json:"receive_time"`
}

func (c *Meters) ToPayload() *payload.MetersRequest {
	metersPayload := &payload.MetersRequest{
		MeterAssetNo: c.MeterAssetNo,
		ReceiveTime:  c.ReceiveTime,
	}
	return metersPayload
}
