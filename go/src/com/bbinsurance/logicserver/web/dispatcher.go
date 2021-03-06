package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/webcommon"
	"net/http"
)

type HandleMethod func(webcommon.BBReq) ([]byte, int, string)

var handlerObjMap map[int]HandleMethod

func FunInitDataBin() {
	handlerObjMap = make(map[int]HandleMethod)
	handlerObjMap[webcommon.FuncListArticle] = FunGetListArticle
	handlerObjMap[webcommon.FuncListCompany] = FunGetListCompany
	handlerObjMap[webcommon.FuncListInsurance] = FunGetListInsurance
	handlerObjMap[webcommon.FuncListComment] = FunGetListComment
	handlerObjMap[webcommon.FuncCreateComment] = FunCreateComment
	handlerObjMap[webcommon.FuncViewComment] = FunViewComment
	handlerObjMap[webcommon.FuncListInsuranceType] = FunGetListInsuranceType
	handlerObjMap[webcommon.FuncGetHomeData] = FunGetHomeData
	handlerObjMap[webcommon.FuncUpComment] = FunUpComment
	handlerObjMap[webcommon.FuncReplyComment] = FunReplyComment
	handlerObjMap[webcommon.FuncGetCompanyById] = FunGetCompanyById
	handlerObjMap[webcommon.FuncGetInsuranceTypeById] = FunGetInsuranceTypeById
	handlerObjMap[webcommon.FuncGetListCommentReply] = FunGetListCommentReply
	handlerObjMap[webcommon.FuncViewArticle] = FunViewArticle
	handlerObjMap[webcommon.FuncGetInsuranceDetail] = FunGetInsuranceDetail
	handlerObjMap[webcommon.FuncListConsultant] = FunListConsultant
}

func FunHandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := webcommon.HandleRequest(request)
	if code != webcommon.ResponseCodeSuccess {
		webcommon.HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		method, ok := handlerObjMap[bbReq.Bin.FunId]
		if ok {
			log.Info("HandleDataBin %d", bbReq.Bin.FunId)
			responseBytes, code, errMsg := method(bbReq)
			if code == webcommon.ResponseCodeSuccess {
				webcommon.HandleSuccessResponse(writer, bbReq, responseBytes)
			} else {
				webcommon.HandleErrorResponse(writer, bbReq, code, errMsg)
			}
		} else {
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeInvalidFunId, "Invalid FunId")
		}
	}
}
