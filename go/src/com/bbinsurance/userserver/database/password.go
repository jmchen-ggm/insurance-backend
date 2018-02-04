package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/userserver/protocol"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const PasswordTableName = "Password"

func InsertPassword(password protocol.Password) protocol.Password {
	sql := fmt.Sprintf("INSERT INTO %s (UserId, PasswordMd5, LastLoginToken, Timestamp) VALUES (?, ?, ?, ?);", PasswordTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		password.UserId = -1
	} else {
		_, err := stmt.Exec(password.UserId, password.PasswordMD5, password.LastLoginToken, password.Timestamp)
		if err != nil {
			password.UserId = -1
		}
	}
	return password
}

func GetPasswordByUserId(userId int64) protocol.Password {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE UserId=?", PasswordTableName)
	rows, err := GetDB().Query(sql, userId)
	defer rows.Close()
	var password protocol.Password
	if err != nil {
		log.Error("GetPasswordByUserId err %s", err)
		password.UserId = -1
	} else {
		if rows.Next() {
			rows.Scan(&password.UserId, &password.PasswordMD5, &password.LastLoginToken, &password.Timestamp)
		} else {
			password.UserId = -1
		}
	}
	return password
}

func UpdateToken(password protocol.Password) protocol.Password {
	sql := fmt.Sprintf("UPDATE %s SET LastLoginToken=?, Timestamp=? WHERE UserId=?", PasswordTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("UpdateToken err %s", err)
		password.UserId = -1
	} else {
		_, err := stmt.Exec(password.LastLoginToken, password.PasswordMD5, password.UserId)
		if err != nil {
			log.Error("UpdateToken err %s", err)
			password.UserId = -1
		}
	}
	return password
}
