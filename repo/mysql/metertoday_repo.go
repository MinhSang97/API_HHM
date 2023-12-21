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

func (s meterTodayRepository) Search(ctx context.Context, meterAssetNo, Start_date, End_date string) ([]model.DataMeter, error) {
	var meters []model.DataMeter

	fmt.Println(Start_date)
	fmt.Println(End_date)

	var meterModel string
	if err := s.db.Raw("SELECT meter_model FROM a_equip_meter WHERE assetno = ?", meterAssetNo).Scan(&meterModel).Error; err != nil {
		return nil, err
	}

	fmt.Println("Meter Model:", meterModel)

	if meterModel == "HHM11-V1" {
		query := `SElECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
            z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
            z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
            FROM
            (SELECT DISTINCT  e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto, e.TV,
            e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
            e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3,
            g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
            g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
            FROM
            (SELECT DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, c.TV,
            FA dn_huucong_giao, FA_T1 dn_huucong_giao_bieu1, FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3,
            FR dn_huucong_nhan, FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
            FROM A_Data_catalogue d
            LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID 
            WHERE d.METER_ASSET_NO = ?
            AND TRUNC(c.tv) >= TO_DATE(?, 'yyyy-mm-dd')
            AND TRUNC(c.tv) <= TO_DATE(?, 'yyyy-mm-dd')) e
            LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID) z
            LEFT join BIZ_PUB_DATA_OTHER_D y
            ON z.Data_ID = y.Data_ID
            and z.nocongto = ? and z.tv = y.tv
            AND TRUNC(y.tv) >= TO_DATE(?, 'yyyy-mm-dd')
    AND TRUNC(y.tv) <= TO_DATE(?, 'yyyy-mm-dd')
    
        AND TRUNC(z.tv) >= TO_DATE(?, 'yyyy-mm-dd')
    AND TRUNC(z.tv) <= TO_DATE(?, 'yyyy-mm-dd')` // Your query for HHM-V1 goes here
		if err := s.db.Raw(query, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, Start_date, End_date).
			Scan(&meters).Error; err != nil {
			fmt.Println(meters)
			return nil, err
		}
	} else if meterModel == "HHM31/38" {
		// Query for HHM31/38
		query := ` SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, 
                            z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
                            z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, 
                            z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
                            z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
                FROM
                    (SELECT DISTINCT e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto, e.tv, 
                            e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
                            e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3, 
                            g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
                            g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
                    FROM
                        (SELECT DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, c.tv, 
                                FA dn_huucong_giao, FA_T1 dn_huucong_giao_bieu1, FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3, 
                                FR dn_huucong_nhan, FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
                        FROM A_Data_catalogue d
                        LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID 
                        WHERE d.METER_ASSET_NO = ?
                AND TRUNC(c.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                            AND TRUNC(c.tv) <= TO_DATE(?, 'yyyy-mm-dd')
                        ) e
                    LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID 
                    WHERE e.METER_ASSET_NO = ?
                        AND g.TV > TO_DATE(?, 'yyyy-mm-dd') 
                        AND g.TV < TO_DATE(?, 'yyyy-mm-dd')
                    ) z
                LEFT JOIN BIZ_PUB_DATA_OTHER_D y
                    ON z.Data_ID = y.Data_ID and z.tv = y.tv
                    WHERE z.nocongto = ?
                    AND TRUNC(y.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                    AND TRUNC(y.tv) <= TO_DATE(?, 'yyyy-mm-dd')
                        AND TRUNC(z.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                    AND TRUNC(z.tv) <= TO_DATE(?, 'yyyy-mm-dd')` // Your query for HHM31/38 goes here
		if err := s.db.Raw(query, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, Start_date, End_date).
			Scan(&meters).Error; err != nil {

			return nil, err
		} else if meterModel == "HHM31" {
			query := ` SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, 
                            z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
                            z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, 
                            z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
                            z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
                FROM
                    (SELECT DISTINCT e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto, e.tv, 
                            e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
                            e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3, 
                            g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
                            g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
                    FROM
                        (SELECT DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, c.tv, 
                                FA dn_huucong_giao, FA_T1 dn_huucong_giao_bieu1, FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3, 
                                FR dn_huucong_nhan, FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
                        FROM A_Data_catalogue d
                        LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID 
                        WHERE d.METER_ASSET_NO = ?
                AND TRUNC(c.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                            AND TRUNC(c.tv) <= TO_DATE(?, 'yyyy-mm-dd')
                        ) e
                    LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID 
                    WHERE e.METER_ASSET_NO = ?
                        AND g.TV > TO_DATE(?, 'yyyy-mm-dd') 
                        AND g.TV < TO_DATE(?, 'yyyy-mm-dd')
                    ) z
                LEFT JOIN BIZ_PUB_DATA_OTHER_D y
                    ON z.Data_ID = y.Data_ID and z.tv = y.tv
                    WHERE z.nocongto = ?
                    AND TRUNC(y.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                    AND TRUNC(y.tv) <= TO_DATE(?, 'yyyy-mm-dd')
                        AND TRUNC(z.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                    AND TRUNC(z.tv) <= TO_DATE(?, 'yyyy-mm-dd')` // Your query for HHM31/38 goes here
			if err := s.db.Raw(query, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, Start_date, End_date).
				Scan(&meters).Error; err != nil {

				return nil, err
			}
		}
	} else if meterModel == "HHM38" {
		// Default query or handle other cases
		query := ` SELECT DISTINCT z.Ma_DIEMDO, z.tenkhachhang, z.nocongto, y.MR_TIME_FA thoigiandoc, 
                            z.dn_huucong_giao, z.dn_huucong_giao_bieu1, z.dn_huucong_giao_bieu2, z.dn_huucong_giao_bieu3,
                            z.dn_huucong_nhan, z.dn_huucong_nhan_bieu1, z.dn_huucong_nhan_bieu2, z.dn_huucong_nhan_bieu3, 
                            z.dn_vocong_giao, z.dn_vocong_giao_bieu1, z.dn_vocong_giao_bieu2, z.dn_vocong_giao_bieu3,
                            z.dn_vocong_nhan, z.dn_vocong_nhan_bieu1, z.dn_vocong_nhan_bieu2, z.dn_vocong_nhan_bieu3
                FROM
                    (SELECT DISTINCT e.Data_ID, e.Ma_DIEMDO, e.tenkhachhang, e.METER_ASSET_NO nocongto, e.tv, 
                            e.dn_huucong_giao, e.dn_huucong_giao_bieu1, e.dn_huucong_giao_bieu2, e.dn_huucong_giao_bieu3,
                            e.dn_huucong_nhan, e.dn_huucong_nhan_bieu1, e.dn_huucong_nhan_bieu2, e.dn_huucong_nhan_bieu3, 
                            g.RA dn_vocong_giao, g.RA_T1 dn_vocong_giao_bieu1, g.RA_T2 dn_vocong_giao_bieu2, g.RA_T3 dn_vocong_giao_bieu3,
                            g.RR dn_vocong_nhan, g.RR_T1 dn_vocong_nhan_bieu1, g.RR_T2 dn_vocong_nhan_bieu2, g.RR_T3 dn_vocong_nhan_bieu3
                    FROM
                        (SELECT DISTINCT d.Data_ID, CONS_NO Ma_DIEMDO, CONS_NAME tenkhachhang, METER_ASSET_NO, c.tv, 
                                FA dn_huucong_giao, FA_T1 dn_huucong_giao_bieu1, FA_T2 dn_huucong_giao_bieu2, FA_T3 dn_huucong_giao_bieu3, 
                                FR dn_huucong_nhan, FR_T1 dn_huucong_nhan_bieu1, FR_T2 dn_huucong_nhan_bieu2, FR_T3 dn_huucong_nhan_bieu3
                        FROM A_Data_catalogue d
                        LEFT JOIN BIZ_PUB_DATA_F_ENERGY_D c ON d.Data_ID = c.Data_ID 
                        WHERE d.METER_ASSET_NO = ?
                AND TRUNC(c.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                            AND TRUNC(c.tv) <= TO_DATE(?, 'yyyy-mm-dd')
                        ) e
                    LEFT JOIN BIZ_PUB_DATA_R_ENERGY_D g ON e.Data_ID = g.Data_ID 
                    WHERE e.METER_ASSET_NO = ?
                        AND g.TV > TO_DATE(?, 'yyyy-mm-dd') 
                        AND g.TV < TO_DATE(?, 'yyyy-mm-dd')
                    ) z
                LEFT JOIN BIZ_PUB_DATA_OTHER_D y
                    ON z.Data_ID = y.Data_ID and z.tv = y.tv
                    WHERE z.nocongto = ?
                    AND TRUNC(y.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                    AND TRUNC(y.tv) <= TO_DATE(?, 'yyyy-mm-dd')
                        AND TRUNC(z.tv) >= TO_DATE(?, 'yyyy-mm-dd')
                    AND TRUNC(z.tv) <= TO_DATE(?, 'yyyy-mm-dd')` // Your query for HHM31/38 goes here
		if err := s.db.Raw(query, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, meterAssetNo, Start_date, End_date, Start_date, End_date).
			Scan(&meters).Error; err != nil {

			return nil, err
		}
	}

	return meters, nil
}

var instancemetertoday meterTodayRepository

func NewMeterTodayRepository(db *gorm.DB) repo.MeterRepoToday {
	if instancemetertoday.db == nil {
		instancemetertoday.db = db

	}
	return instancemetertoday
}
