package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/userserver/protocol"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const UserTableName = "User"

func InsertUser(user protocol.User) protocol.User {
	sql := fmt.Sprintf("INSERT INTO %s (Username, Nickname, Timestamp, ThumbUrl) VALUES (?, ?, ?, ?);", UserTableName)
	log.Info("InsertUser sql=%s", sql)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		user.Id = -1
	} else {
		result, err := stmt.Exec(user.Username, user.Nickname, user.Timestamp, user.ThumbUrl)
		if err != nil {
			user.Id = -1
		} else {
			user.Id, _ = result.LastInsertId()
		}
	}
	return user
}

func ListUser(startIndex int, length int) []protocol.User {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s", UserTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", UserTableName, length, startIndex)
	}
	log.Info("ListUser sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var userList []protocol.User
	if err != nil {
		log.Error("ListUser err %s", err)
	} else {
		for rows.Next() {
			var user protocol.User
			rows.Scan(&user.Id, &user.Username, &user.Nickname, &user.Timestamp, &user.ThumbUrl)
			userList = append(userList, user)
		}
		log.Info("ListUser %d ", len(userList))
	}
	return userList
}

func GetUserById(id int64) protocol.User {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id=?", UserTableName)
	rows, err := GetDB().Query(sql, id)
	defer rows.Close()
	var user protocol.User
	user.Id = -1
	if err != nil {
		log.Error("GetUser err %s", err)
	} else {
		if rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Nickname, &user.Timestamp, &user.ThumbUrl)
		}
	}
	return user
}

func GetUserByUsername(username string) protocol.User {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE Username=?", UserTableName)
	rows, err := GetDB().Query(sql, username)
	defer rows.Close()
	var user protocol.User
	user.Id = -1
	if err != nil {
		log.Error("GetUser err %s", err)
	} else {
		if rows.Next() {
			rows.Scan(&user.Id, &user.Username, &user.Nickname, &user.Timestamp, &user.ThumbUrl)
		}
	}
	return user
}
