package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"com/bbinsurance/util"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const CommentTableName = "Comment"

func InsertComment(comment Comment) (int64, error) {
	log.Info("InsertComment %s", util.objToString(comment))
	sql := fmt.Sprintf("INSERT INTO %s (Uin, Content, TotalScore, Score1, Score2, Score3, Timestamp, ViewCount, Flags) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);", CommentTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		timestamp := time.GetTimestamp()
		result, err := stmt.Exec(comment.Uin, comment.Content, comment.TotalScore, comment.Score1, comment.Score2, comment.Score3, timestamp, 0, protocol.CREATED)
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
		sql = fmt.Sprintf("SELECT * FROM %s ORDER BY Timestamp DESC;", CommentTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s ORDER BY Timestamp DESC LIMIT %d OFFSET %d;", CommentTableName, length, startIndex)
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
			rows.Scan(&comment.ServerId, &comment.Uin, &comment.Content, &comment.TotalScore, &comment.Score1, &comment.Score2, &comment.Score3, &comment.Timestamp, &comment.ViewCount, &comment.Flags)
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
