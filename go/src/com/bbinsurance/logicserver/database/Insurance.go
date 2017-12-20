package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const InsuranceTableName = "Insurance"

func InsertInsurance(nameZHCN string, nameEN string, desc string, InsuranceTypeId int64, companyId int, thumbUrl string) (int64, error) {
	sql := fmt.Sprintf("INSERT INTO %s (NameZHCN, NameEN, Desc, Type, Timestamp, CompanyId, ThumbUrl) VALUES (?, ?, ?, ?, ?, ?, ?);", InsuranceTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		timestamp := time.GetTimestampInMilli()
		result, err := stmt.Exec(nameZHCN, nameEN, desc, InsuranceTypeId, timestamp, companyId, thumbUrl)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return -1, err
		} else {
			id, err := result.LastInsertId()
			return id, err
		}
	}
}

func UpdateInsuranceThumbUrl(id int64, thumbUrl string) {
	log.Info("UpdateInsuranceThumbUrl: id=%d thumbUrl=%s", id, thumbUrl)
	sql := fmt.Sprintf("UPDATE %s SET thumbUrl=? WHERE id= ?;", InsuranceTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(thumbUrl, id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
		} else {
			log.Info("UpdateInsuranceThumbUrl Success")
		}
	}
}

func GetListInsurance(startIndex int, length int) []protocol.Insurance {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf(
			"SELECT %s.Id,%s.NameZHCN,%s.NameEN,%s.Desc,%s.Name,%s.Name,%s.Timestamp,%s.ThumbUrl FROM %s,%s,%s where %s.CompanyId=%s.Id and %s.Type=%s.Id",
			InsuranceTableName, InsuranceTableName, InsuranceTableName, InsuranceTableName, InsuranceTypeTableName, CompanyTableName, InsuranceTableName, InsuranceTableName,
			InsuranceTableName, CompanyTableName, InsuranceTypeTableName, InsuranceTableName, CompanyTableName, InsuranceTableName, InsuranceTypeTableName)
	} else {
		sql = fmt.Sprintf(
			"SELECT %s.Id,%s.NameZHCN,%s.NameEN,%s.Desc,%s.Name,%s.Name,%s.Timestamp,%s.ThumbUrl FROM %s,%s,%s where %s.CompanyId=%s.Id and %s.Type=%s.Id LIMIT %d OFFSET %d",
			InsuranceTableName, InsuranceTableName, InsuranceTableName, InsuranceTableName, InsuranceTypeTableName, CompanyTableName, InsuranceTableName, InsuranceTableName,
			InsuranceTableName, CompanyTableName, InsuranceTypeTableName, InsuranceTableName, CompanyTableName, InsuranceTableName, InsuranceTypeTableName, length, startIndex)
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
			rows.Scan(&insurance.Id, &insurance.NameZHCN, &insurance.NameEN, &insurance.Desc, &insurance.Type, &insurance.Company, &insurance.Timestamp, &insurance.ThumbUrl)
			insuranceList = append(insuranceList, insurance)
		}
		log.Info("GetListInsurance %d ", len(insuranceList))
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
