package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/util"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const CommentReplyTableName = "CommentReply"

func InsertCommentReply(commentReply protocol.CommentReply) protocol.CommentReply {
	commentReply.Id = -1
	log.Info("InsertCommentReply %s", util.ObjToString(commentReply))
	sql := fmt.Sprintf("INSERT INTO %s (Uin, ReplyUin, CommentId, Content, Timestamp) VALUES (?, ?, ?, ?, ?);", CommentReplyTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return commentReply
	} else {
		commentReply.Timestamp = 
		result, err := stmt.Exec(commentReply.Uin, commentReply.ReplyUin, commentReply.CommentId, commentReply.Content, commentReply.Timestamp)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			commentReply.Id = -1
		} else {
			commentReply.Id, _ = result.LastInsertId()
		}
	}
	return commentReply
}

func GetCommentReplyListByCommentId(commentId int64) []protocol.CommentReply {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE CommentId = ? ORDER BY Timestamp DESC", CommentReplyTableName)
	rows, err := GetDB().Query(sql, commentId)
	defer rows.Close()
	var commentReplyList []protocol.CommentReply
	if err != nil {
		log.Error("GetReplyCommentListByCommentId err %s", err)
	} else {
		for rows.Next() {
			var commentReply protocol.CommentReply
			rows.Scan(&commentReply.Id, &commentReply.Uin, &commentReply.ReplyUin,
				&commentReply.CommentId, &commentReply.Content, &commentReply.Timestamp)
			commentReplyList = append(commentReplyList, commentReply)
		}
	}
	return commentReplyList
}
