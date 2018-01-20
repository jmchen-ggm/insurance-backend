package database

import (
	"com/bbinsurance/log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	log.Info("InitDB Start")
	var err error
	db, err = sql.Open("sqlite3", "./kv.db")
	if err != nil {
		log.Error("open db err = %s", err)
	} else {
		log.Info("open db success")
	}

	var createHttpRequestSql = "CREATE TABLE IF NOT EXISTS HttpFun(Id INTEGER PRIMARY KEY AUTOINCREMENT, FunId INTEGER, Timestamp INTEGER, ResponseSize INTEGER, UseTime INTEGER, Uin INTEGER);"
	_, err = db.Exec(createHttpRequestSql, nil)
	if err != nil {
		log.Error("Create HttpRequest Error: sql = %s, err = %s", createHttpRequestSql, err)
	} else {
		log.Info("Create HttpRequest Table Success sql = %s", createHttpRequestSql)
	}
}

func GetDB() *sql.DB {
	return db
}
