package service

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
)

func GetListConsultant(startIndex int, length int) []protocol.Consultant {
	consultantList := database.GetListConsultant(startIndex, length)
	return consultantList
}
