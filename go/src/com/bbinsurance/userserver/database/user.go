package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/time"
	"com/bbinsurance/userserver/protocol"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const UserTableName = "User"

func InsertUser(user protocol.User) (int64, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Username, Nickname, PhoneNumber, Timestamp, ThumbUrl) VALUES (?, ?, ?, ?, ?);", UserTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		timestamp := time.GetTimestampInMilli()
		result, err := stmt.Exec(user.Username, user.Nickname, user.PhoneNumber, timestamp, user.ThumbUrl)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return -1, err
		} else {
			id, err := result.LastInsertId()
			return id, err
		}
	}
}

func GetUser(id int64) (protocol.User, error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", UserTableName)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var user protocol.User
	user.Id = -1
	if err != nil {
		log.Error("GetUser err %s", err)
		return user, err
	} else {
		if rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Nickname, &user.PhoneNumber, &user.Timestamp, &user.ThumbUrl)
			return user, nil
		} else {
			return user, errors.New("Not Found User")
		}
	}
}

func GetUserByUsername(username string) (protocol.User, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE Username=?", UserTableName)
	rows, err := GetDB().Query(sql, username)
	defer rows.Close()
	var user protocol.User
	user.Id = -1
	if err != nil {
		log.Error("GetUser err %s", err)
		return user, err
	} else {
		if rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Nickname, &user.PhoneNumber, &user.Timestamp, &user.ThumbUrl)
			return user, nil
		} else {
			return user, errors.New("Not Found User")
		}
	}
}
