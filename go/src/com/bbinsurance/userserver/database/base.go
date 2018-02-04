package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/util"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const UserTableVersionKey = "UserTableVersionKey"
const PasswordTableVersionKey = "PasswordTableVersionKey"

const CurrentUserTableVersion = "1"
const CurrentPasswordTableVersion = "1"

var db *sql.DB

func InitDB() {
	log.Info("InitDB Start")
	var err error
	db, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Error("open db err = %s", err)
	} else {
		log.Info("open db success")
	}

	util.CreateDBConfigTable(db)
	util.CheckDBTable(db, UserTableName, UserTableVersionKey, CurrentUserTableVersion)
	util.CheckDBTable(db, PasswordTableName, PasswordTableVersionKey, CurrentPasswordTableVersion)

	var createUserSql = "CREATE TABLE IF NOT EXISTS User(Id INTEGER PRIMARY KEY AUTOINCREMENT, Username TEXT NOT NULL, NickName TEXT NOT NULL, Timestamp INTEGER, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createUserSql, nil)
	if err != nil {
		log.Error("Create User Error: sql = %s, err = %s", createUserSql, err)
	} else {
		log.Info("Create User Table Success sql = %s", createUserSql)
	}

	var createUsernameIndexSql = "CREATE INDEX IF NOT EXISTS User_Username On User(Username)"
	db.Exec(createUsernameIndexSql, nil)

	var createPasswordSql = "CREATE TABLE IF NOT EXISTS Password(UserId INTEGER PRIMARY KEY, PasswordMd5 TEXT NOT NULL, LastLoginToken TEXT NOT NULL, Timestamp INTEGER);"
	_, err = db.Exec(createPasswordSql, nil)
	if err != nil {
		log.Error("Create Password Error: sql = %s, err = %s", createPasswordSql, err)
	} else {
		log.Info("Create Password Table Success sql = %s", createPasswordSql)
	}

	util.SetSequenceStartId(db, UserTableName, 10000)
}

func GetDB() *sql.DB {
	return db
}
