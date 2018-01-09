package service

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
)

var insuranceTypeCacheMap map[int64]protocol.InsuranceType

func GetInsuranceTypeById(id int64) protocol.InsuranceType {
	if insuranceTypeCacheMap == nil {
		insuranceTypeCacheMap = make(map[int64]protocol.InsuranceType)
	}
	insuranceType, ok := insuranceTypeCacheMap[id]
	if ok {
		log.Info("Hit InsuranceType Service Cache %d", id)
		return insuranceType
	} else {
		insuranceType = database.GetInsuranceTypeById(id)
		if insuranceType.Id != -1 {
			insuranceTypeCacheMap[id] = insuranceType
		}
		return insuranceType
	}
}

func GetListInsurance(startIndex int, length int) []protocol.Insurance {
	insuranceList := database.GetListInsurance(startIndex, length)
	insuranceListLength := len(insuranceList)
	for i := 0; i < insuranceListLength; i++ {
		insuranceList[i].CompanyName = GetCompanyById(insuranceList[i].CompanyId).Name
		insuranceList[i].InsuranceTypeName = GetInsuranceTypeById(insuranceList[i].InsuranceTypeId).Name
	}
	return insuranceList
}