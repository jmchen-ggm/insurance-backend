package service

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
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
