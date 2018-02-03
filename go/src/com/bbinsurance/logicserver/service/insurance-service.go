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
		insuranceList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(insuranceList[i].ThumbUrl)
	}
	return insuranceList
}

func GetListInsuranceType(startIndex int, length int) []protocol.InsuranceType {
	insuranceTypeList := database.GetListInsuranceType(startIndex, length)
	insuranceTypeListLength := len(insuranceTypeList)
	for i := 0; i < insuranceTypeListLength; i++ {
		insuranceTypeList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(insuranceTypeList[i].ThumbUrl)
	}
	return insuranceTypeList
}

func GetInsuranceDetail(id int64) protocol.InsuranceDetail {
	insurance := database.GetInsuranceById(id)
	var insuranceDetail protocol.InsuranceDetail
	if insurance.Id >= 0 {
		insuranceDetail.Id = insurance.Id
		insuranceDetail.Name = insurance.Name
		insuranceDetail.Desc = insurance.Desc
		insuranceDetail.InsuranceType = GetInsuranceTypeById(insurance.InsuranceTypeId)
		insuranceDetail.Company = GetCompanyById(insurance.CompanyId)
		insuranceDetail.AgeFrom = insurance.AgeFrom
		insuranceDetail.AgeTo = insurance.AgeTo
		insuranceDetail.AnnualCompensation = insurance.AnnualCompensation
		insuranceDetail.AnnualPremium = insurance.AnnualPremium
		insuranceDetail.Flags = insurance.Flags
		insuranceDetail.Timestamp = insurance.Timestamp
		insuranceDetail.ThumbUrl = insurance.ThumbUrl
		insuranceDetail.DetailData = insurance.DetailData
	} else {
		insuranceDetail.Id = -1
	}
	return insuranceDetail
}
