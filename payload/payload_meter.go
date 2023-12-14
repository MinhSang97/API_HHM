package payload

import (
	"app/model"
	"encoding/json"
	"log"
)

type MetersRequest struct {
	MeterAssetNo int64  `json:"meter_asset_no"`
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
