package payload

import (
	"app/model"
	"encoding/json"
	"log"
)

type MetersRequest struct {
	MeterAssetNo string `json:"meter_asset_no"`
	ReceiveTime  string `json:"receive_time"`
}

func (c *MetersRequest) ToModel() *model.Meters {
	meter := &model.Meters{
		MeterAssetNo: c.MeterAssetNo,
		ReceiveTime:  c.ReceiveTime,
	}

	return meter
}

func (c *MetersRequest) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
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

func (c *AddMeterRequest) ToModel() *model.DataMeter {
	meter := &model.DataMeter{
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
