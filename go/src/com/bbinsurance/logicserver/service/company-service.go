package service

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
)

var companyCacheMap map[int64]protocol.Company

func GetCompanyById(id int64) protocol.Company {
	if companyCacheMap == nil {
		companyCacheMap = make(map[int64]protocol.Company)
	}
	company, ok := companyCacheMap[id]
	if ok {
		log.Info("Hit Company Service Cache %d", id)
		return company
	} else {
		company = database.GetCompanyById(id)
		if company.Id != -1 {
			companyCacheMap[id] = company
		}
		return company
	}
}

func GetListCompany(startIndex int, length int) []protocol.Company {
	companyList := database.GetListCompany(startIndex, length)
	companyListLength := len(companyList)
	for i := 0; i < companyListLength; i++ {
		companyList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(companyList[i].ThumbUrl)
	}
	return companyList
}
