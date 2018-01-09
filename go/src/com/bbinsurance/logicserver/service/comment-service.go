package service

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
)

func GetListComment(startIndex int, length int) []protocol.Comment {
	commentList := database.GetListComment(startIndex, length)
	commentListLength := len(commentList)
	for i := 0; i < commentListLength; i++ {
		commentList[i].CompanyName = GetCompanyById(commentList[i].CompanyId).Name
		commentList[i].InsuranceTypeName = GetInsuranceTypeById(commentList[i].InsuranceTypeId).Name
	}
	return commentList
}

func UpComment(commentUp protocol.CommentUp, isUp bool) protocol.Comment {
	database.UpdateCommentUpCount(commentUp.CommentId, isUp)
	if isUp {
		database.InsertCommentUp(commentUp)
	} else {
		database.DeleteCommentUp(commentUp.Uin, commentUp.CommentId)
	}
	comment := database.GetCommentById(commentUp.CommentId)
	return comment
}

func ReplyComment(commentReply protocol.CommentReply) protocol.Comment {
	database.UpdateCommentReplyCount(commentReply.CommentId)
	database.InsertCommentReply(commentReply)
	comment := database.GetCommentById(commentReply.CommentId)
	return comment
}

func ViewComment(id int64) protocol.Comment {
	database.UpdateCommentViewCount(id)
	comment := database.GetCommentById(id)
	return comment
}
