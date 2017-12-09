package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/userserver/protocol"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const PasswordTableName = "Password"

func InsertPassword(password protocol.Password) error {
	sql := fmt.Sprintf("INSERT INTO %s (UserId, PasswordMd5, LastLoginToken, Timestamp) VALUES (?, ?, ?, ?);", PasswordTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return err
	} else {
		_, err := stmt.Exec(password.UserId, password.PasswordMD5, password.LastLoginToken, password.Timestamp)
		return err
	}
}

func GetPasswordByUserId(userId int64) (protocol.Password, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE UserId=?", PasswordTableName)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var password protocol.Password
	if err != nil {
		log.Error("GetUser err %s", err)
		return password, err
	} else {
		if rows.Next() {
			rows.Scan(&password.UserId, &password.PasswordMD5, &password.LastLoginToken, &password.Timestamp)
			return password, nil
		} else {
			return password, errors.New("Not Found Password")
		}
	}
}

func UpdateToken(password protocol.Password) error {
	sql := fmt.Sprintf("UPDATE %s SET LastLogicToken=?, Timestamp=? WHERE UserId=?", PasswordTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return err
	} else {
		_, err := stmt.Exec(password.LastLoginToken, password.PasswordMD5, password.UserId)
		return err
	}
}
