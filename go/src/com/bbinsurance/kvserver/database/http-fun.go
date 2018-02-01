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

func ListHttpFunByTime(beginTime int64, endTime int64) []webcommon.HttpFun {
	sql := fmt.Sprint("SELECT * FROM %s WHERE Timestamp >= %d AND endTime <= %d", HttpFunTableName, beginTime, endTime)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var httpFunList []webcommon.HttpFun
	if err != nil {
		log.Error("ListHttpFun err %s", err)
		return httpFunList
	} else {
		for rows.Next() {
			var httpFun webcommon.HttpFun
			rows.Scan(&httpFun.Id, &httpFun.FunId, &httpFun.Timestamp, &httpFun.ResponseSize, &user.UseTime, &user.Uin)
			httpFunList = append(httpFunList, httpFun)
		}
		log.Info("ListHttpFun %d ", len(httpFunList))
		return httpFunList
	}
}

func ListHttpFunByPage(startIndex int, pageSize int) []webcommon.HttpFun {
	sql := fmt.Sprint("SELECT * FROM %s WHERE Timestamp >= %d AND endTime <= %d", HttpFunTableName, beginTime, endTime)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var httpFunList []webcommon.HttpFun
	if err != nil {
		log.Error("ListHttpFun err %s", err)
		return httpFunList
	} else {
		for rows.Next() {
			var httpFun webcommon.HttpFun
			rows.Scan(&httpFun.Id, &httpFun.FunId, &httpFun.Timestamp, &httpFun.ResponseSize, &user.UseTime, &user.Uin)
			httpFunList = append(httpFunList, httpFun)
		}
		log.Info("ListHttpFun %d ", len(httpFunList))
		return httpFunList
	}
}
