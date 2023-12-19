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
	MeterAssetNo string `json:"meter_asset_no"`
	ReceiveTime  string `json:"receive_time"`
}

func (c *Meters) ToPayload() *payload.MetersRequest {
	metersPayload := &payload.MetersRequest{
		MeterAssetNo: c.MeterAssetNo,
		ReceiveTime:  c.ReceiveTime,
	}
	return metersPayload
}

type MetersRequestToday struct {
	MeterAssetNo string `json:"meter_asset_no"`
	Start_date   string `json:"start_date"`
	End_date     string `json:"end_date"`
}

func (c *MetersRequestToday) ToPayload() *payload.MetersRequestToday {
	meter := &payload.MetersRequestToday{
		MeterAssetNo: c.MeterAssetNo,
		Start_date:   c.Start_date,
		End_date:     c.End_date,
	}

	return meter
}

type AddMeterRequest struct {
	MA_DIEMDO             string
	TENKHACHHANG          string
	NOCONGTO              string `json:"meter_asset_no"`
	THOIGIANDOC           string `json:"receive_time"`
	DN_HUUCONG_GIAO       float64
	DN_HUUCONG_GIAO_BIEU1 float64
	DN_HUUCONG_GIAO_BIEU2 float64
	DN_HUUCONG_GIAO_BIEU3 float64
	DN_HUUCONG_NHAN       float64
	DN_HUUCONG_NHAN_BIEU1 float64
	DN_HUUCONG_NHAN_BIEU2 float64
	DN_HUUCONG_NHAN_BIEU3 float64
	DN_VOCONG_GIAO        float64
	DN_VOCONG_GIAO_BIEU1  float64
	DN_VOCONG_GIAO_BIEU2  float64
	DN_VOCONG_GIAO_BIEU3  float64
	DN_VOCONG_NHAN        float64
	DN_VOCONG_NHAN_BIEU1  float64
	DN_VOCONG_NHAN_BIEU2  float64
	DN_VOCONG_NHAN_BIEU3  float64
}

func (c *AddMeterRequest) ToPayload() *payload.AddMeterRequest {
	meter := &payload.AddMeterRequest{
		MA_DIEMDO:             c.MA_DIEMDO,
		TENKHACHHANG:          c.TENKHACHHANG,
		NOCONGTO:              c.THOIGIANDOC,
		THOIGIANDOC:           c.THOIGIANDOC,
		DN_HUUCONG_GIAO:       c.DN_HUUCONG_GIAO,
		DN_HUUCONG_GIAO_BIEU1: c.DN_HUUCONG_GIAO_BIEU1,
		DN_HUUCONG_GIAO_BIEU2: c.DN_HUUCONG_GIAO_BIEU2,
		DN_HUUCONG_GIAO_BIEU3: c.DN_HUUCONG_GIAO_BIEU3,
		DN_HUUCONG_NHAN:       c.DN_HUUCONG_NHAN,
		DN_HUUCONG_NHAN_BIEU1: c.DN_HUUCONG_NHAN_BIEU1,
		DN_HUUCONG_NHAN_BIEU2: c.DN_HUUCONG_NHAN_BIEU2,
		DN_HUUCONG_NHAN_BIEU3: c.DN_HUUCONG_NHAN_BIEU3,
		DN_VOCONG_GIAO:        c.DN_VOCONG_GIAO,
		DN_VOCONG_GIAO_BIEU1:  c.DN_VOCONG_GIAO_BIEU1,
		DN_VOCONG_GIAO_BIEU2:  c.DN_VOCONG_GIAO_BIEU2,
		DN_VOCONG_GIAO_BIEU3:  c.DN_VOCONG_GIAO_BIEU3,
		DN_VOCONG_NHAN:        c.DN_VOCONG_NHAN,
		DN_VOCONG_NHAN_BIEU1:  c.DN_VOCONG_NHAN_BIEU1,
		DN_VOCONG_NHAN_BIEU2:  c.DN_VOCONG_GIAO_BIEU2,
		DN_VOCONG_NHAN_BIEU3:  c.DN_VOCONG_NHAN_BIEU3,
	}

	return meter
}
