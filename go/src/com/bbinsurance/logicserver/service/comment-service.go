package service

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
)

func GetListComment(uin int64, startIndex int, length int) []protocol.Comment {
	commentList := database.GetListComment(startIndex, length)
	commentListLength := len(commentList)
	for i := 0; i < commentListLength; i++ {
		commentList[i].CompanyName = GetCompanyById(commentList[i].CompanyId).Name
		commentList[i].InsuranceTypeName = GetInsuranceTypeById(commentList[i].InsuranceTypeId).Name
		commentList[i].IsUp = database.CheckCommentUp(uin, commentList[i].Id)
	}
	return commentList
}

func UpComment(commentUp protocol.CommentUp, isUp bool) protocol.Comment {
	dbUp := database.CheckCommentUp(commentUp.Uin, commentUp.CommentId)
	comment := database.GetCommentById(commentUp.CommentId)
	log.Info("database.CheckCommentUp dbUp: %t", dbUp)
	if dbUp == isUp {
		comment.IsUp = isUp
		return comment
	} else {
		var canUpdateCount = false
		if isUp {
			if database.InsertCommentUp(commentUp).Id >= 0 {
				canUpdateCount = true
			}
		} else {
			canUpdateCount = database.DeleteCommentUp(commentUp.Uin, commentUp.CommentId)
		}
		if canUpdateCount {
			if isUp {
				comment.UpCount++
			} else {
				comment.UpCount--
			}
			database.UpdateCommentUpCount(commentUp.CommentId, comment.UpCount)
			comment.IsUp = isUp
		} else {
			comment.IsUp = dbUp
		}
		return comment
	}
}

func ReplyComment(uin int64, commentReply protocol.CommentReply) protocol.Comment {
	database.UpdateCommentReplyCount(commentReply.CommentId)
	database.InsertCommentReply(commentReply)
	comment := database.GetCommentById(commentReply.CommentId)
	comment.IsUp = database.CheckCommentUp(uin, comment.Id)
	return comment
}

func ViewComment(uin int64, id int64) protocol.Comment {
	database.UpdateCommentViewCount(id)
	comment := database.GetCommentById(id)
	comment.IsUp = database.CheckCommentUp(uin, id)
	return comment
}
