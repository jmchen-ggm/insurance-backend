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

func GetListCommentReply(commentId int64) []protocol.CommentReply {
	return database.GetReplyCommentListByCommentId(commentId)
}

func UpComment(commentUp protocol.CommentUp, isUp bool) protocol.Comment {
	dbUp := database.CheckCommentUp(commentUp.Uin, commentUp.CommentId)
	comment := GetCommentById(commentUp.Uin, commentUp.CommentId)
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
	comment := GetCommentById(uin, commentReply.CommentId)
	comment.IsUp = database.CheckCommentUp(uin, commentReply.CommentId)
	commentReply = database.InsertCommentReply(commentReply)
	if commentReply.Id >= 0 {
		comment.ReplyCount++
		database.UpdateCommentReplyCount(commentReply.CommentId, comment.ReplyCount)
	}
	return comment
}

func ViewComment(uin int64, id int64) protocol.Comment {
	comment := GetCommentById(uin, id)
	comment.IsUp = database.CheckCommentUp(uin, id)
	comment.ViewCount++
	database.UpdateCommentViewCount(id, comment.ViewCount)
	return comment
}

func GetCommentById(uin int64, id int64) protocol.Comment {
	comment := database.GetCommentById(id)
	comment.CompanyName = GetCompanyById(comment.CompanyId).Name
	comment.InsuranceTypeName = GetInsuranceTypeById(comment.InsuranceTypeId).Name
	comment.IsUp = database.CheckCommentUp(uin, comment.Id)
	return comment
}
