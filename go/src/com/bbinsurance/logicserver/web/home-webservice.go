package web

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
	"encoding/json"
)

func FunGetHomeData(bbReq webcommon.BBReq) ([]byte, int, string) {
	var response protocol.BBGetHomeDataResponse
	response.BannerList = database.GetTopBannerInsuranceList()
	response.TopCommentList = database.GetTopCommentList()
	response.TopInsuranceTypeList = database.GetListInsuranceType(0, -1)
	response.TopCompanyList = database.GetListCompany(0, -1)
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}
