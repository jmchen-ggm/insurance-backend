package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const InsuranceTypeTableName = "InsuranceType"

func InsertType(name string) (int64, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Name) VALUES (?);", InsuranceTypeTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		result, err := stmt.Exec(name)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return -1, err
		} else {
			id, err := result.LastInsertId()
			return id, err
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
	log.Info("GetListCompany sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var typeList []protocol.Type
	if err != nil {
		log.Error("GetListCompany err %s", err)
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
	sql = fmt.Sprintf("SELECT * FROM %s where Name=? limit 1",InsuranceTypeTableName)
	log.Info("Get Company By Name sql=%s", sql)
	rows, err := GetDB().Query(sql,Name)
	defer rows.Close()
	if err != nil {
		log.Error("GetCompany err %s", err)
	} else {
		for rows.Next() {
			err = rows.Scan(&Type.Id, &Type.Name)
			if err != nil {
				log.Error("GetCompany err %s", err)
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
			log.Info("RemoveCompanyById %d Success", id)
		}
	}
}
