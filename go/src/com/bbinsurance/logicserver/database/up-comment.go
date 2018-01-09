package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/util"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const CommentUpTableName = "CommentUp"

func InsertCommentUp(commentUp protocol.CommentUp) protocol.CommentUp {
	commentUp.Id = -1
	log.Info("InsertCommentUp %s", util.ObjToString(commentUp))
	sql := fmt.Sprintf("INSERT INTO %s (Uin, CommentId, Timestamp) VALUES (?, ?, ?);", CommentUpTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return commentUp
	} else {
		result, err := stmt.Exec(commentUp.Uin, commentUp.CommentId, commentUp.Timestamp)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			commentUp.Id = -1
		} else {
			commentUp.Id, err = result.LastInsertId()
		}
	}
	return commentUp
}

func CheckCommentUp(uin int64, commentId int64) bool {
	sql := fmt.Sprintf("SELECT 1 FROM %s WHERE Uin = ? AND CommentId = ?", CommentUpTableName)
	rows, err := GetDB().Query(sql, uin, commentId)
	defer rows.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return false
	} else {
		return rows.Next()
	}
}

func DeleteCommentUp(uin int64, commentId int64) bool {
	sql := fmt.Sprintf("DELETE FROM %s WHERE Uin = ? AND CommentId = ?", CommentUpTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return false
	} else {
		result, err := stmt.Exec(uin, commentId)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return false
		} else {
			rowCnt, err := result.RowsAffected()
			if rowCnt > 0 {
				log.Info("DeleteCommentUp uin:%d commentId:%d Success", uin, commentId)
				return true
			} else {
				return false
			}
		}
	}
}
