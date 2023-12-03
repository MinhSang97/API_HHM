package model

import "time"

type Login struct {
	user     string
	password string
}

type Meter struct{
	meter_asset_so int64
	thoi_gian time.Time
}

func (c *Login) TableName() string {
	return "login"
}
