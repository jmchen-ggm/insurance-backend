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
	db, err = sql.Open("sqlite3", "./logic.db")
	if err != nil {
		log.Error("open db err = %s", err)
	} else {
		log.Info("open db success")
	}
	var createArticleSql = "CREATE TABLE IF NOT EXISTS Article(Id INTEGER PRIMARY KEY AUTOINCREMENT, Title Text NOT NULL, Desc Text NOT NULL, Url Text NOT NULL, ThumbUrl Text NOT NULL)"
	_, err = db.Exec(createArticleSql, nil)
	if err != nil {
		log.Error("Create Article Error: sql = %s, err = %s", createArticleSql, err)
	} else {
		log.Info("Create Article Table Success sql = %s", createArticleSql)
	}
}

func GetDB() *sql.DB {
	return db
}
