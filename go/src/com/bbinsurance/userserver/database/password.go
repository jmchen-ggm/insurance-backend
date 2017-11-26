package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/time"
	"com/bbinsurance/userserver/protocol"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InsertPassword(password protocol.Password) error {
	sql := fmt.Sprintf("INSERT INTO %s (UserId, PasswordMd5, LastLoginToken, Timestamp) VALUES (?, ?, ?, ?, ?);", UserTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		timestamp := time.GetTimestampInMilli()
		_, err := stmt.Exec(password.UserId, password.PasswordMD5, password.LastLoginToken, password.Timestamp)
		return err
	}
}
