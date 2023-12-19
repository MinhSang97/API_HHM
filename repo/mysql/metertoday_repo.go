package mysql

import (
	"app/model"
	"app/repo"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type meterTodayRepository struct {
	db *gorm.DB
}

func (s meterTodayRepository) Search(ctx context.Context, meterAssetNo, start_date, end_date string) ([]model.DataMeter, error) {
	var meters []model.DataMeter

	start_date = "%" + start_date + "%"

	end_date = "%" + end_date + "%"

	var meterModel string
	if err := s.db.Raw("SELECT meter_model FROM a_equip_meter WHERE assetno = ?", meterAssetNo).Scan(&meterModel).Error; err != nil {
		return nil, err
	}

	fmt.Println("Meter Model:", meterModel)

	if meterModel == "HHM-V1" {
		// Query for HHM-V1
		query := `SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
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
	LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID WHERE d.METER_ASSET_NO = ? AND c.RECEIVE_TIME LIKE ?) e
	LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID) z
	LEFT join BIZ_PUB_DATA_OTHER_D y
	ON z.Data_ID = y.Data_ID and y.MR_TIME_FA LIKE ?` // Your query for HHM-V1 goes here
		if err := s.db.Raw(query, meterAssetNo, start_date, end_date).
			Scan(&meters).Error; err != nil {

			return nil, err
		}
	} else if meterModel == "HHM31/38" {
		// Query for HHM31/38
		query := `SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
            z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
            z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
            FROM
            (SElECT DISTINCT  e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto,e.TV thoigiandoc, e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
            e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3, g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
            g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
            FROM
            (SELECT  DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, c.TV, FA dn_huucong_giao,FA_T1 dn_huucong_giao_bieu1,
            FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3, FR dn_huucong_nhan,FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
            FROM A_Data_catalogue d
            LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c 
            ON d.Data_ID = c.Data_ID where d.METER_ASSET_NO = ? and c.RECEIVE_TIME LIKE ?)e
            LEFT join BIZ_PUB_DATA_R_ENERGY_D g
            ON e.Data_ID = g.Data_ID where e.METER_ASSET_NO = ? and g.tv LIKE ?)z
            LEFT join BIZ_PUB_DATA_OTHER_D y
            ON z.Data_ID = y.Data_ID and y.MR_TIME_FA LIKE ?` // Your query for HHM31/38 goes here
		if err := s.db.Raw(query, meterAssetNo, end_date, meterAssetNo, end_date, end_date).
			Scan(&meters).Error; err != nil {

			return nil, err
		} else if meterModel == "HHM31" {
			query := `SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
            z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
            z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
            FROM
            (SElECT DISTINCT  e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto,e.TV thoigiandoc, e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
            e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3, g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
            g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
            FROM
            (SELECT  DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, c.TV, FA dn_huucong_giao,FA_T1 dn_huucong_giao_bieu1,
            FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3, FR dn_huucong_nhan,FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
            FROM A_Data_catalogue d
            LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c 
            ON d.Data_ID = c.Data_ID where d.METER_ASSET_NO = ? and c.RECEIVE_TIME LIKE ?)e
            LEFT join BIZ_PUB_DATA_R_ENERGY_D g
            ON e.Data_ID = g.Data_ID where e.METER_ASSET_NO = ? and g.tv LIKE ?)z
            LEFT join BIZ_PUB_DATA_OTHER_D y
            ON z.Data_ID = y.Data_ID and y.MR_TIME_FA LIKE ?` // Your query for HHM31/38 goes here
			if err := s.db.Raw(query, meterAssetNo, end_date, meterAssetNo, end_date, end_date).
				Scan(&meters).Error; err != nil {

				return nil, err
			}
		}
	} else if meterModel == "HHM38" {
		// Default query or handle other cases
		query := `SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
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
	LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID WHERE d.METER_ASSET_NO = ? AND c.RECEIVE_TIME LIKE ?) e
	LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID) z
	LEFT join BIZ_PUB_DATA_OTHER_D y
	ON z.Data_ID = y.Data_ID and y.MR_TIME_FA LIKE ?` // Your default query goes here
		if err := s.db.Raw(query, meterAssetNo, end_date, end_date).
			Scan(&meters).Error; err != nil {

			return nil, err
		}
	}

	// do something with result

	return meters, nil
}

var instancemetertoday meterTodayRepository

func NewMeterTodayRepository(db *gorm.DB) repo.MeterRepoToday {
	if instancemetertoday.db == nil {
		instancemetertoday.db = db

	}
	return instancemetertoday
}
