package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/util"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const CommentTableName = "Comment"

func InsertComment(comment protocol.Comment) (protocol.Comment, error) {
	log.Info("InsertComment %s", util.ObjToString(comment))
	sql := fmt.Sprintf("INSERT INTO %s (Uin, Content, CompanyId, InsuranceTypeId, TotalScore, Score1, Score2, Score3, Score4, Timestamp, UpCount, ViewCount, ReplyCount, Flags) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", CommentTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		comment.Id = -1
	} else {
		result, err := stmt.Exec(comment.Uin, comment.Content, comment.CompanyId, comment.InsuranceTypeId, comment.TotalScore, comment.Score1,
			comment.Score2, comment.Score3, comment.Score4, comment.Timestamp, comment.UpCount, comment.ViewCount, comment.ReplyCount, comment.Flags)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			comment.Id = -1
		} else {
			comment.Id, err = result.LastInsertId()
		}
	}
	return comment, err
}

func UpdateCommentViewCount(id int64, viewCount int) {
	sql := fmt.Sprintf("UPDATE %s SET ViewCount=%d WHERE id=?;", CommentTableName, viewCount)
	log.Info("UpdateCommentViewCount: id=%d sql=%s", id, sql)
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

func UpdateCommentUpCount(id int64, upCount int) {
	sql := fmt.Sprintf("UPDATE %s SET UpCount=%d WHERE Id=?;", CommentTableName, upCount)
	log.Info("UpdateCommentUpCount: id=%d sql=%s", id, sql)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
		} else {
			log.Info("UpdateCommentUpCount Success")
		}
	}
}

func UpdateCommentReplyCount(id int64, replyCount int) {
	sql := fmt.Sprintf("UPDATE %s SET ReplyCount=%d WHERE Id = ?", CommentTableName, replyCount)
	log.Info("UpdateCommentReplyCount: id=%d sql=%s", id, sql)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
		} else {
			log.Info("UpdateCommentReplyCount Success")
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
			rows.Scan(&comment.Id, &comment.Uin, &comment.Content, &comment.CompanyId, &comment.InsuranceTypeId,
				&comment.TotalScore, &comment.Score1, &comment.Score2, &comment.Score3, &comment.Score4,
				&comment.Timestamp, &comment.UpCount, &comment.ViewCount, &comment.ReplyCount, &comment.Flags)
			commentList = append(commentList, comment)
		}
		log.Info("GetListComment %d ", len(commentList))
	}
	return commentList
}

func GetTopCommentList() []protocol.Comment {
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY UpCount DESC LIMIT 3;", CommentTableName)
	log.Info("GetTopComment sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var commentList []protocol.Comment
	if err != nil {
		log.Error("GetTopComment err %s", err)
	} else {
		for rows.Next() {
			var comment protocol.Comment
			rows.Scan(&comment.Id, &comment.Uin, &comment.Content, &comment.CompanyId, &comment.InsuranceTypeId,
				&comment.TotalScore, &comment.Score1, &comment.Score2, &comment.Score3, &comment.Score4,
				&comment.Timestamp, &comment.UpCount, &comment.ViewCount, &comment.ReplyCount, &comment.Flags)
			commentList = append(commentList, comment)
		}
		log.Info("GetTopComment %d ", len(commentList))
	}
	return commentList
}

func GetCommentById(id int64) protocol.Comment {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE Id = ?", CommentTableName)
	log.Info("GetCommentById sql=%s", sql)
	rows, err := GetDB().Query(sql, id)
	defer rows.Close()
	var comment protocol.Comment
	comment.Id = -1
	if err != nil {
		log.Error("GetCommentById err %s", err)
		return comment
	} else {
		if rows.Next() {
			rows.Scan(&comment.Id, &comment.Uin, &comment.Content, &comment.CompanyId, &comment.InsuranceTypeId,
				&comment.TotalScore, &comment.Score1, &comment.Score2, &comment.Score3, &comment.Score4,
				&comment.Timestamp, &comment.UpCount, &comment.ViewCount, &comment.ReplyCount, &comment.Flags)
			return comment
		} else {
			return comment
		}
	}
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
			log.Info("DeleteCommentById %d Success", id)
		}
	}
}
