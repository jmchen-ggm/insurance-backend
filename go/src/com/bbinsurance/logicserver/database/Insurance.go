package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const InsuranceTableName = "Insurance"
const InsuranceListField = "Id, Name, Desc, InsuranceTypeId, CompanyId, AgeFrom, AgeTo, AnnualCompensation, AnnualPremium, Flags, Timestamp, ThumbUrl"

func InsertInsurance(insurance protocol.Insurance) (protocol.Insurance, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Name, Desc, InsuranceTypeId, CompanyId, AgeFrom, AgeTo, AnnualCompensation, AnnualPremium, Flags, Timestamp, ThumbUrl, DetailData) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", InsuranceTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		insurance.Id = -1
	} else {
		insurance.Timestamp = time.GetTimestampInMilli()
		result, err := stmt.Exec(insurance.Name, insurance.Desc, insurance.InsuranceTypeId, insurance.CompanyId,
			insurance.AgeFrom, insurance.AgeTo, insurance.AnnualCompensation, insurance.AnnualPremium,
			insurance.Flags, insurance.Timestamp, insurance.ThumbUrl, insurance.DetailData)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			insurance.Id = -1
		} else {
			insurance.Id, err = result.LastInsertId()
		}
	}
	return insurance, err
}

func GetInsuranceById(id int64) protocol.Insurance {
	sql := fmt.Sprintf("SELECT %s, DetailData FROM %s WHERE Id = ?;", InsuranceListField, InsuranceTableName)
	log.Info("GetInsuranceById sql=%s", sql)
	rows, err := GetDB().Query(sql, id)
	defer rows.Close()
	var insurance protocol.Insurance
	if err != nil {
		insurance.Id = -1
		log.Error("GetListInsurance err %s", err)
	} else {
		if rows.Next() {
			rows.Scan(&insurance.Id, &insurance.Name, &insurance.Desc, &insurance.InsuranceTypeId,
				&insurance.CompanyId, &insurance.AgeFrom, &insurance.AgeTo, &insurance.AnnualCompensation,
				&insurance.AnnualPremium, &insurance.Flags, &insurance.Timestamp, &insurance.ThumbUrl, &insurance.DetailData)
		} else {
			insurance.Id = -1
		}
	}
	return insurance
}

func GetListInsurance(startIndex int, length int) []protocol.Insurance {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT %s FROM %s;", InsuranceListField, InsuranceTableName)
	} else {
		sql = fmt.Sprintf("SELECT %s FROM %s LIMIT %d OFFSET %d", InsuranceListField, InsuranceTableName, length, startIndex)
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
			rows.Scan(&insurance.Id, &insurance.Name, &insurance.Desc, &insurance.InsuranceTypeId,
				&insurance.CompanyId, &insurance.AgeFrom, &insurance.AgeTo, &insurance.AnnualCompensation,
				&insurance.AnnualPremium, &insurance.Flags, &insurance.Timestamp, &insurance.ThumbUrl)
			insuranceList = append(insuranceList, insurance)
		}
		log.Info("GetListInsurance %d ", len(insuranceList))
	}
	return insuranceList
}

func GetTopBannerInsuranceList() []protocol.Insurance {
	sql := fmt.Sprintf("SELECT %s FROM %s LIMIT 5", InsuranceListField, InsuranceTableName)
	log.Info("GetTopBannerInsuranceList sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var insuranceList []protocol.Insurance
	if err != nil {
		log.Error("GetTopBannerInsuranceList err %s", err)
	} else {
		for rows.Next() {
			var insurance protocol.Insurance
			rows.Scan(&insurance.Id, &insurance.Name, &insurance.Desc, &insurance.InsuranceTypeId,
				&insurance.CompanyId, &insurance.AgeFrom, &insurance.AgeTo, &insurance.AnnualCompensation,
				&insurance.AnnualPremium, &insurance.Flags, &insurance.Timestamp, &insurance.ThumbUrl)
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
