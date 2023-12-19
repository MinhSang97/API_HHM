package mysql

import (
	"app/model"
	"app/repo"
	"context"
	"gorm.io/gorm"
)

type meterRepository struct {
	db *gorm.DB
}

func (s meterRepository) Search(ctx context.Context, MeterAssetNo, ReceiveTime string) ([]model.DataMeter, error) {
	var meters []model.DataMeter

	query := `SElECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
		z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
		z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
	FROM
	(SELECT DISTINCT  e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto, e.TV thoigiandoc,
		e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
		e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3,
		g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
		g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
	FROM
	(SELECT DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, TV,
		FA dn_huucong_giao, FA_T1 dn_huucong_giao_bieu1, FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3,
		FR dn_huucong_nhan, FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
	FROM A_Data_catalogue d
	LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID WHERE d.METER_ASSET_NO = :MeterAssetNo AND c.RECEIVE_TIME LIKE :ReceiveTime) e
	LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID)z
	LEFT join BIZ_PUB_DATA_OTHER_D y
	ON z.Data_ID = y.Data_ID and y.MR_TIME_FA LIKE :ReceiveTime`

	if err := s.db.Raw(query, gorm.Expr("MeterAssetNo = ? AND ReceiveTime LIKE ?", MeterAssetNo, ReceiveTime)).
		Scan(&meters).Error; err != nil {

		return nil, err
	}

	// do something with result

	return meters, nil
}

var instancemeter meterRepository

func NewMeterRepository(db *gorm.DB) repo.MeterRepo {
	if instancemeter.db == nil {
		instancemeter.db = db

	}
	return instancemeter
}
