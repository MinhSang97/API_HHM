package model

import (
	"time"
)

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

type MetersToday struct {
	MeterAssetNo string `json:"meter_asset_no"`
	Start_date   string `json:"start_date"`
	End_date     string `json:"end_date"`
}

func (c *MetersToday) TableName() string {
	return "meterstoday"
}

type DataMeter struct {
	MA_DIEMDO             string
	TENKHACHHANG          string
	NOCONGTO              string
	THOIGIANDOC           string
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
