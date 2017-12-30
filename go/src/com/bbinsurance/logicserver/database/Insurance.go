package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const InsuranceTableName = "Insurance"

func InsertInsurance(insurance protocol.Insurance) (protocol.Insurance, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Name, Desc, InsuranceTypeId, CompanyId, Timestamp, ThumbUrl, DetailData) VALUES (?, ?, ?, ?, ?, ?, ?);", InsuranceTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		insurance.Id = -1
	} else {
		insurance.Timestamp = time.GetTimestampInMilli()
		result, err := stmt.Exec(insurance.Name, insurance.Desc, insurance.InsuranceTypeId, insurance.CompanyId,
			insurance.Timestamp, insurance.ThumbUrl, insurance.DetailData)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			insurance.Id = -1
		} else {
			insurance.Id, err = result.LastInsertId()
		}
	}
	return insurance, err
}

func GetListInsurance(startIndex int, length int) []protocol.Insurance {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf(
			"SELECT A.Id, A.Name, A.Desc, B.Id, B.Name, C.Id, C.Name, A.Timestamp, A.ThumbUrl, A.DetailData FROM %s AS A, %s AS B, %s AS C where CompanyId=C.Id and InsuranceTypeId=B.Id",
			InsuranceTableName, InsuranceTypeTableName, CompanyTableName)
	} else {
		sql = fmt.Sprintf(
			"SELECT A.Id, A.Name, A.Desc, B.Id, B.Name, C.Id, C.Name, A.Timestamp, A.ThumbUrl, A.DetailData FROM %s AS A, %s AS B, %s AS C where CompanyId=C.Id and InsuranceTypeId=B.Id LIMIT %d OFFSET %d",
			InsuranceTableName, InsuranceTypeTableName, CompanyTableName, length, startIndex)
	}
	log.Info("GetListInsurance sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var insuranceList []protocol.Insurance
	if err != nil {
		log.Error("GetListInsurance err %s", err)
	} else {
		for rows.Next() {
			var insurance protocol.Insurance
			rows.Scan(&insurance.Id, &insurance.Name, &insurance.Desc, &insurance.InsuranceTypeId, &insurance.InsuranceTypeName,
				&insurance.CompanyId, &insurance.CompanyName, &insurance.Timestamp, &insurance.ThumbUrl, &insurance.DetailData)
			insuranceList = append(insuranceList, insurance)
		}
		log.Info("GetListInsurance %d ", len(insuranceList))
	}
	return insuranceList
}

func GetTopBannerInsuranceList() []protocol.Insurance {
	sql := fmt.Sprintf(
		"SELECT A.Id, A.Name, A.Desc, B.Id, B.Name, C.Id, C.Name, A.Timestamp, A.ThumbUrl, A.DetailData FROM %s AS A, %s AS B, %s AS C WHERE CompanyId=C.Id and InsuranceTypeId=B.Id ORDER BY Timestamp DESC LIMIT 5",
		InsuranceTableName, InsuranceTypeTableName, CompanyTableName)
	log.Info("GetTopBannerInsuranceList sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var insuranceList []protocol.Insurance
	if err != nil {
		log.Error("GetTopBannerInsuranceList err %s", err)
	} else {
		for rows.Next() {
			var insurance protocol.Insurance
			rows.Scan(&insurance.Id, &insurance.Name, &insurance.Desc, &insurance.InsuranceTypeId, &insurance.InsuranceTypeName,
				&insurance.CompanyId, &insurance.CompanyName, &insurance.Timestamp, &insurance.ThumbUrl, &insurance.DetailData)
			insuranceList = append(insuranceList, insurance)
		}
		log.Info("GetTopBannerInsuranceList %d ", len(insuranceList))
	}
	return insuranceList
}

func DeleteInsuranceById(id int64) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", InsuranceTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return
		} else {
			log.Info("RemoveInsuranceById %d Success", id)
		}
	}
}
