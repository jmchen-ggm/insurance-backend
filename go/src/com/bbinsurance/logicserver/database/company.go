package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func InsertCompany(name string, desc string, thumbUrl string) (int64, error) {
	stmt, err := GetDB().Prepare("INSERT INTO Company (Name, Desc, ThumbUrl) VALUES (?, ?, ?);")
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	}

	result, err := stmt.Exec(name, desc, thumbUrl)
	if err != nil {
		log.Error("Prepare Exec Error %s", err)
		return -1, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func UpdateCompanyThumbUrl(id int64, thumbUrl string) {
	log.Info("UpdateCompanyThumbUrl: id=%d thumbUrl=%s", id, thumbUrl)
	stmt, err := GetDB().Prepare("UPDATE Company SET thumbUrl=? WHERE id=?;")
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return
	}

	_, err = stmt.Exec(thumbUrl, id)
	if err != nil {
		log.Error("Prepare Exec Error %s", err)
		return
	} else {
		log.Info("UpdateCompanyThumbUrl Success")
	}
}

func GetListCompany(startIndex int, length int) []protocol.Company {
	var sql string
	if (length == -1) {
		sql = fmt.Sprintf("SELECT * FROM Company")
	} else {
		sql = fmt.Sprintf("SELECT * FROM Company LIMIT %d OFFSET %d", length, startIndex)
	}
	log.Info("GetListCompany sql=%s", sql)
	rows, err := GetDB().Query(sql)
	if err != nil {
		log.Error("GetListCompany err %s", err)
	}
	var companyList []protocol.Company
	for rows.Next() {
		var company protocol.Company
		rows.Scan(&company.Id, &company.Name, &company.Desc, &company.ThumbUrl)
		companyList = append(companyList, company)
	}
	rows.Close()
	return companyList
}