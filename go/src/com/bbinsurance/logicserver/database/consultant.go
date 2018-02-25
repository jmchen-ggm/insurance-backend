package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const ConsultantTableName = "Consultant"

func InsertConsultant(consultant protocol.Consultant) (protocol.Consultant, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Name, Desc, Score, ThumbUrl, Flags, DetailData) VALUES (?, ?, ?, ?, ?, ?);", ConsultantTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		consultant.Id = -1
	} else {
		result, err := stmt.Exec(consultant.Name, consultant.Desc, consultant.Score, consultant.ThumbUrl, consultant.Flags, consultant.DetailData)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			consultant.Id = -1
		} else {
			consultant.Id, err = result.LastInsertId()
		}
	}
	return consultant, err
}

func GetListConsultant(startIndex int, length int) []protocol.Consultant {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s", ConsultantTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", ConsultantTableName, length, startIndex)
	}
	log.Info("GetListConsultant sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var consultantList []protocol.Consultant
	if err != nil {
		log.Error("GetListConsultant err %s", err)
	} else {
		for rows.Next() {
			var consultant protocol.Consultant
			rows.Scan(&consultant.Id, &consultant.Name, &consultant.Desc, &consultant.Score, &consultant.ThumbUrl, &consultant.Flags, &consultant.DetailData)
			consultantList = append(consultantList, consultant)
		}
		log.Info("GetListConsultant %d ", len(consultantList))
	}
	return consultantList
}
