package payload

import (
	"app/model"
	"time"
)

type AddTokensRequest struct {
	Token      string    `json:"token"`
	UserName   string    `json:"user_name" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Logindate  time.Time `json:"login_date"`
	ExpireDate time.Time `json:"expire_date"`
	ExpireTime time.Time `json:"expire_time"`
	IPAddress  string    `json:"ip_address"`
}

func (c *AddTokensRequest) ToModel() *model.Tokens {
	token := &model.Tokens{
		Token:      c.Token,
		UserName:   c.UserName,
		Password:   c.Password,
		LoginDate:  c.Logindate,
		ExpireDate: c.ExpireTime,
		ExpireTime: c.ExpireTime,
		IPAddress:  c.IPAddress,
	}

	return token
}
