package service

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
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

func GetListInsuranceType(startIndex int, length int) []protocol.InsuranceType {
	insuranceTypeList := database.GetListInsuranceType(startIndex, length)
	insuranceListLength := len(insuranceTypeList)
	for i := 0; i < insuranceListLength; i++ {
		insuranceTypeList[i].CompanyName = GetCompanyById(insuranceTypeList[i].CompanyId).Name
		insuranceTypeList[i].InsuranceTypeName = GetInsuranceTypeById(insuranceTypeList[i].InsuranceTypeId).Name
		insuranceTypeList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(insuranceTypeList[i].ThumbUrl)
	}
	return insuranceTypeList
}
