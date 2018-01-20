package service

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
)

func GetHomeData() protocol.BBGetHomeDataResponse {
	var response protocol.BBGetHomeDataResponse
	response.BannerList = database.GetTopBannerInsuranceList()
	for i := 0; i < len(response.BannerList); i++ {
		response.BannerList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(response.BannerList[i].ThumbUrl)
	}
	response.TopCommentList = database.GetTopCommentList()
	response.TopInsuranceTypeList = database.GetListInsuranceType(0, -1)
	for i := 0; i < len(response.TopInsuranceTypeList); i++ {
		response.TopInsuranceTypeList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(response.TopInsuranceTypeList[i].ThumbUrl)
	}
	response.TopCompanyList = database.GetListCompany(0, -1)
	for i := 0; i < len(response.TopCompanyList); i++ {
		response.TopCompanyList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(response.TopCompanyList[i].ThumbUrl)
	}
	return response
}
