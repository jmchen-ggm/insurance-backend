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
	var createArticleSql = "CREATE TABLE IF NOT EXISTS Article(Id INTEGER PRIMARY KEY AUTOINCREMENT, Title TEXT NOT NULL, Desc TEXT NOT NULL, Date TEXT NOT NULL, TimeStamp INTEGER, Url TEXT NOT NULL, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createArticleSql, nil)
	if err != nil {
		log.Error("Create Article Error: sql = %s, err = %s", createArticleSql, err)
	} else {
		log.Info("Create Article Table Success sql = %s", createArticleSql)
	}
	var createCompanySql = "CREATE TABLE IF NOT EXISTS Company(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT Not NULL, Desc TEXT NOT NULL, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createCompanySql, nil)
	if err != nil {
		log.Error("Create Company Error: sql = %s, err = %s", createCompanySql, err)
	} else {
		log.Info("Create Company Table Success sql = %s", createCompanySql)
	}
	var createInsuranceSql = "CREATE TABLE IF NOT EXISTS Insurance(Id INTEGER PRIMARY KEY AUTOINCREMENT, NameZHCN TEXT Not NULL, NameEN TEXT NOT NULL, Desc TEXT NOT NULL, Type INTEGER, CompanyId INTEGER, Timestamp INTEGER, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createInsuranceSql, nil)
	if err != nil {
		log.Error("Create Insurance Error: sql = %s, err = %s", createInsuranceSql, err)
	} else {
		log.Info("Create Insurance Table Success sql = %s", createInsuranceSql)
	}
	var createCommentSql = "CREATE TABLE IF NOT EXISTS Comment(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, Content TEXT NOT NULL, Score INTEGER, TimeStamp INTEGER, ViewCount INTEGER, Flags INTEGER);"
	_, err = db.Exec(createCommentSql, nil)
	if err != nil {
		log.Error("Create Comment Error: sql = %s, err = %s", createCommentSql, err)
	} else {
		log.Info("Create Comment Table Success sql = %s", createCommentSql)
	}

	var createCommentTimestampIndex = "CREATE INDEX IF NOT EXISTS Comment_Timestamp ON Comment(Timestamp)"
	db.Exec(createCommentTimestampIndex, nil)

	var createSubCommentSql = "CREATE TABLE IF NOT EXISTS SubComment(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, ReplyUin INTEGER, CommentId INTEGER, Content TEXT NOT NULL, TimeStamp INTEGER);"
	_, err = db.Exec(createSubCommentSql, nil)
	if err != nil {
		log.Error("Create SubComment Error: sql = %s, err = %s", createSubCommentSql, err)
	} else {
		log.Info("Create SubComment Table Success sql = %s", createSubCommentSql)
	}
}

func GetDB() *sql.DB {
	return db
}
