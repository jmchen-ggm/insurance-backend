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

func GetListInsuranceType(startIndex int, length int) []protocol.Type {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s", InsuranceTypeTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", InsuranceTypeTableName, length, startIndex)
	}
	log.Info("GetListInsuranceType sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var typeList []protocol.Type
	if err != nil {
		log.Error("GetListInsuranceType err %s", err)
	} else {
		for rows.Next() {
			var Type protocol.Type
			rows.Scan(&Type.Id, &Type.Name)
			typeList = append(typeList, Type)
		}
		log.Info("GetListType %d ", len(typeList))
	}
	return typeList
}

func SelectInsuranceTypeByName(Name string) protocol.Type {
	var sql string
	var Type protocol.Type
	sql = fmt.Sprintf("SELECT * FROM %s where Name=? limit 1", InsuranceTypeTableName)
	log.Info("Get InsuranceType By Name sql=%s", sql)
	rows, err := GetDB().Query(sql, Name)
	defer rows.Close()
	if err != nil {
		log.Error("GetInsuranceType err %s", err)
	} else {
		for rows.Next() {
			err = rows.Scan(&Type.Id, &Type.Name)
			if err != nil {
				log.Error("GetInsuranceType err %s", err)
			}
		}
	}
	return Type
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
