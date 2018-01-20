package main

import (
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/util"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	HandleInsuranceType()
}

func HandleInsuranceType() {
	insuranceTypeJsonStr, _ := util.FileGetContent("/Users/jiaminchen/develop/insurance/insurance-file/static/data/insurance-types.json")
	fmt.Printf("%s\n", insuranceTypeJsonStr)
	var listInsuranceTypeResponse protocol.BBListInsuranceTypeResponse
	json.Unmarshal(util.StringToBytes(insuranceTypeJsonStr), &listInsuranceTypeResponse)
	fmt.Printf("%s\n", util.ObjToString(listInsuranceTypeResponse))
	insuranceTypeList := listInsuranceTypeResponse.InsuranceTypeList
	for i := 0; i < len(insuranceTypeList); i++ {
		InsertInsuranceType(insuranceTypeList[i])
	}
}

func InsertInsuranceType(insuranceType protocol.InsuranceType) (protocol.InsuranceType, error) {
	db, _ := sql.Open("sqlite3", "./logic.db")
	sql := fmt.Sprintf("INSERT INTO InsuranceType (Id, Name, Desc, ThumbUrl, Flags, DetailData) VALUES (?, ?, ?, ?, ?, ?);")
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		insuranceType.Id = -1
	} else {
		result, err := stmt.Exec(insuranceType.Id, insuranceType.Name, insuranceType.Desc, insuranceType.ThumbUrl, insuranceType.Flags, insuranceType.DetailData)
		if err != nil {
			insuranceType.Id = -1
		} else {
			insuranceType.Id, err = result.LastInsertId()
		}
	}
	return insuranceType, err
}
