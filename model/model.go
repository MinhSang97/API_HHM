package model

import (
	"encoding/json"
	"log"
	"time"
)

type Student struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age"`
	Grade        float32   `json:"grade"`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"entrance_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *Student) TableName() string {
	return "students"
}

type Tokens struct {
	Token      string    `json:"token"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	LoginDate  time.Time `json:"login_date"`
	ExpireDate time.Time `json:"expire_Date"`
	ExpireTime time.Time `json:"expire_time"`
	IPAddress  string    `json:"ip_address"`
}

func (c *Tokens) TableName() string {
	return "tokens"
}

type Meters struct {
	MeterAssetNo string `json:"meter_asset_no"`
	ReceiveTime  string `json:"receive_time"`
}

func (c *Meters) TableName() string {
	return "meters"
}

type DataMeter struct {
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

func (c *DataMeter) TableName() string {
	return "datameters"
}

func (c *Student) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *Student) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
