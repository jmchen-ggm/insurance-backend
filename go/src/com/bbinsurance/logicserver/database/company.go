package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const CompanyTableName = "Company"

func InsertCompany(company protocol.Company) (protocol.Company, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Name, Desc, ThumbUrl) VALUES (?, ?, ?);", CompanyTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		company.Id = -1
	} else {
		result, err := stmt.Exec(company.Name, company.Desc, company.ThumbUrl)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			company.Id = -1
		} else {
			company.Id, err = result.LastInsertId()
		}
	}
	return company, err
}

func UpdateCompanyThumbUrl(id int64, thumbUrl string) {
	log.Info("UpdateCompanyThumbUrl: id=%d thumbUrl=%s", id, thumbUrl)
	sql := fmt.Sprintf("UPDATE %s SET thumbUrl=? WHERE id= ?;", CompanyTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(thumbUrl, id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
		} else {
			log.Info("UpdateCompanyThumbUrl Success")
		}
	}
}

func GetCompanyNameById(id int64) string {
	sql := fmt.Sprintf("SELECT name FROM %s WHERE id = ?", CompanyTableName)
	rows, err := GetDB().Query(sql, id)
	if err != nil {
		log.Error("GetCompanyNameById err %s", err)
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

func GetListCompany(startIndex int, length int) []protocol.Company {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s", CompanyTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", CompanyTableName, length, startIndex)
	}
	log.Info("GetListCompany sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var companyList []protocol.Company
	if err != nil {
		log.Error("GetListCompany err %s", err)
	} else {
		for rows.Next() {
			var company protocol.Company
			rows.Scan(&company.Id, &company.Name, &company.Desc, &company.ThumbUrl)
			companyList = append(companyList, company)
		}
		log.Info("GetListCompany %d ", len(companyList))
	}
	return companyList
}

func SelectCompanyByName(Name string) protocol.Company {
	var sql string
	var company protocol.Company
	sql = fmt.Sprintf("SELECT * FROM %s where Name=?limit 1", CompanyTableName)
	log.Info("Get Company By Name sql=%s", sql)
	rows, err := GetDB().Query(sql, Name)
	defer rows.Close()
	if err != nil {
		log.Error("GetCompany err %s", err)
	} else {
		for rows.Next() {
			rows.Scan(&company.Id, &company.Name, &company.Desc, &company.ThumbUrl)
		}
		log.Info("GetCompany %s ", company.Name)
	}
	return company
}

func DeleteCompanyById(id int64) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE Id=?", CompanyTableName)
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
