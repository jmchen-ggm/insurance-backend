package service

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
)

func GetListConsultant(startIndex int, length int) []protocol.Consultant {
	consultantList := database.GetListConsultant(startIndex, length)
	consultantListLength = len(consultantList)
	for i := 0; i < consultantListLength; i++ {
		consultantList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(consultantList[i].ThumbUrl)
	}
	return consultantList
}
