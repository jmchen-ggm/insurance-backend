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
		log.Info("Hit InsuranceType Service Cache %s", id)
		return insuranceType
	} else {
		insuranceType = database.GetInsuranceTypeById(id)
		if insuranceType.Id != -1 {
			insuranceTypeCacheMap[id] = insuranceType
		}
		return insuranceType
	}
}
