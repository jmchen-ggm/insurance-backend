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
	sql := fmt.Sprintf("INSERT INTO %s (Uin, ReplyUin CommentId, Content, Timestamp) VALUES (?, ?, ?, ?, ?);", CommentTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return commentReply
	} else {
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
