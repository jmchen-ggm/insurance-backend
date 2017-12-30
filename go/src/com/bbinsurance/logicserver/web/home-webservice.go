package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/constants"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
	"encoding/json"
	"fmt"
)

func FunGetHomeData(bbReq webcommon.BBReq) {
	var response BBGetHomeDataResponse
	response.BannerList = database.GetTopBannerInsuranceList()
	response.TopComment = database.GetTopComment()
	response.TopInsuranceTypeList = database.GetListInsuranceType(0, -1)
	response.TopCompanyList = database.GetListCompany(0, -1)
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}
