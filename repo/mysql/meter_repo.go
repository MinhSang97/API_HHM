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

type DataCatalogue struct {
	gorm.Model
	Data_ID               string
	Ma_DIEMDO             string
	Tenkhachhang          string
	Meter_Asset_No        string
	TV                    string
	Dn_huucong_giao       float64
	Dn_huucong_giao_bieu1 float64
	Dn_huucong_giao_bieu2 float64
	Dn_huucong_giao_bieu3 float64
	Dn_huucong_nhan       float64
	Dn_huucong_nhan_bieu1 float64
	Dn_huucong_nhan_bieu2 float64
	Dn_huucong_nhan_bieu3 float64
}

type EnergyData struct {
	gorm.Model
	Data_ID              string
	Dn_vocong_giao       float64
	Dn_vocong_giao_bieu1 float64
	Dn_vocong_giao_bieu2 float64
	Dn_vocong_giao_bieu3 float64
	Dn_vocong_nhan       float64
	Dn_vocong_nhan_bieu1 float64
	Dn_vocong_nhan_bieu2 float64
	Dn_vocong_nhan_bieu3 float64
}

type OtherData struct {
	gorm.Model
	Data_ID    string
	MR_Time_FA string
}

func (s meterRepository) Search(ctx context.Context, MeterAssetNo, ReceiveTime string) ([]model.Meters, error) {
	var meters []model.Meters

	// Use ? instead of % in the WHERE clause
	if err := s.db.Table("A_Data_Catalogue").Select("DISTINCT Data_ID, Ma_DIEMDO, Tenkhachhang, Meter_Asset_No AS Nocongto, TV AS Thoigiandoc").
		Joins("LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D ON A_Data_Catalogue.Data_ID = BIZ_PUB_DATA_F_ENERGY_D.Data_ID AND A_Data_Catalogue.Meter_Asset_No = ? AND BIZ_PUB_DATA_F_ENERGY_D.RECEIVE_TIME LIKE ?", MeterAssetNo, ReceiveTime).
		Joins("LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D ON A_Data_Catalogue.Data_ID = BIZ_PUB_DATA_R_ENERGY_D.Data_ID AND BIZ_PUB_DATA_R_ENERGY_D.TV LIKE ?", ReceiveTime).
		Joins("LEFT JOIN BIZ_PUB_DATA_OTHER_D ON A_Data_Catalogue.Data_ID = BIZ_PUB_DATA_OTHER_D.Data_ID AND BIZ_PUB_DATA_OTHER_D.MR_Time_FA LIKE ?", ReceiveTime).
		Scan(&meters).Error; err != nil {
		return nil, err
	}

	return meters, nil
}

var instancemeter meterRepository

func NewMeterRepository(db *gorm.DB) repo.MeterRepo {
	if instancemeter.db == nil {
		instancemeter.db = db

	}
	return instancemeter
}
