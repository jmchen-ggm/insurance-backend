package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const CommentTableName = "Comment"

func InsertComment(uin int64, context string, score int) (int64, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Uin, Content, Score, Timestamp, ViewCount, Flags) VALUES (?, ?, ?, ?, ?, ?);", CommentTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		timestamp := time.GetTimestamp()
		result, err := stmt.Exec(uin, context, score, timestamp, 0, protocol.CREATED)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return -1, err
		} else {
			id, err := result.LastInsertId()
			return id, err
		}
	}
}

func UpdateCommentViewCount(id int64) {
	log.Info("UpdateCommentViewCount: id=%d", id)
	sql := fmt.Sprintf("UPDATE %s SET ViewCount=ViewCount+1 WHERE id=?;", CommentTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
		} else {
			log.Info("UpdateCommentViewCount Success")
		}
	}
}

func GetListComment(startIndex int, length int) []protocol.Comment {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s ORDER BY TimeStamp DESC;", CommentTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s ORDER BY TimeStamp DESC LIMIT %d OFFSET %d;", CommentTableName, length, startIndex)
	}
	log.Info("GetListComment sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var commentList []protocol.Comment
	if err != nil {
		log.Error("GetListComment err %s", err)
	} else {
		for rows.Next() {
			var comment protocol.Comment
			rows.Scan(&comment.ServId, &comment.Uin, &comment.Content, &comment.Score, &comment.Timestamp, &comment.ViewCount, &comment.Flags)
			commentList = append(commentList, comment)
		}
		log.Info("GetListComment %d ", len(commentList))
	}
	return commentList
}

func DeleteCommentById(id int64) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", CommentTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return
		} else {
			log.Info("RemoveCommentById %d Success", id)
		}
	}
}
