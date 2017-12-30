package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const InsuranceTypeTableName = "InsuranceType"

func InsertInsuranceType(insuranceType protocol.InsuranceType) (protocol.InsuranceType, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Name, Desc, ThumbUrl) VALUES (?, ?, ?);", InsuranceTypeTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		insuranceType.Id = -1
	} else {
		result, err := stmt.Exec(insuranceType.Name, insuranceType.Desc, insuranceType.ThumbUrl)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			insuranceType.Id = -1
		} else {
			insuranceType.Id, err = result.LastInsertId()
		}
	}
	return insuranceType, err
}

func GetInsuranceTypeNameById(id int64) string {
	sql := fmt.Sprintf("SELECT name FROM %s WHERE id = ?", InsuranceTypeTableName)
	rows, err := GetDB().Query(sql, id)
	if err != nil {
		log.Error("GetInsuranceTypeNameById error %s", err)
		return ""
	} else {
		if rows.Next() {
			var name string
			rows.Scan(&name)
			return name
		} else {
			return ""
		}
	}
}

func GetListInsuranceType(startIndex int, length int) []protocol.InsuranceType {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s", InsuranceTypeTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", InsuranceTypeTableName, length, startIndex)
	}
	log.Info("GetListInsuranceType sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var insuranceTypeList []protocol.InsuranceType
	if err != nil {
		log.Error("GetListInsuranceType err %s", err)
	} else {
		for rows.Next() {
			var insuranceType protocol.InsuranceType
			rows.Scan(&insuranceType.Id, &insuranceType.Name, &insuranceType.Desc, &insuranceType.ThumbUrl)
			insuranceTypeList = append(insuranceTypeList, insuranceType)
		}
		log.Info("GetListInsuranceType %d ", len(insuranceTypeList))
	}
	return insuranceTypeList
}

func DeleteInsuranceTypeById(id int64) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE Id=?", InsuranceTypeTableName)
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
			log.Info("RemoveInsuranceTypeById %d Success", id)
		}
	}
}
