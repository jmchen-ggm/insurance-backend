package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/webcommon"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const HttpFunTableName = "HttpFun"

func InsertHttpFun(httpFun webcommon.HttpFun) webcommon.HttpFun {
	sql := fmt.Sprintf("INSERT INTO %s (FunId, Timestamp, ResponseSize, UseTime, Uin) VALUES (?, ?, ?, ?, ?);", HttpFunTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		httpFun.Id = -1
		return httpFun
	} else {
		result, err := stmt.Exec(httpFun.FunId, httpFun.Timestamp, httpFun.ResponseSize, httpFun.UseTime, httpFun.Uin)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			httpFun.Id = -1
			return httpFun
		} else {
			httpFun.Id, err = result.LastInsertId()
			return httpFun
		}
	}
}
