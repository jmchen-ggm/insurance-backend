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
