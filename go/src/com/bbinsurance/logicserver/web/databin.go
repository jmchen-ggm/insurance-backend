package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/webcommon"
	"net/http"
)

type HandleMethod func(webcommon.BBReq) ([]byte, int, string)

var dandlerObjMap map[int]HandleMethod

func FunInitDataBin() {
	dandlerObjMap = make(map[int]HandleMethod)
	dandlerObjMap[webcommon.FuncListArticle] = FunGetListArticle
	dandlerObjMap[webcommon.FuncListCompany] = FunGetListCompany
	dandlerObjMap[webcommon.FuncListInsurance] = FunGetListInsurance
	dandlerObjMap[webcommon.FuncListComment] = FunGetListComment
	dandlerObjMap[webcommon.FuncCreateComment] = FunCreateComment
	dandlerObjMap[webcommon.FuncViewComment] = FunViewComment
}

func FunHandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := HandleRequest(request)
	if code != webcommon.ResponseCodeSuccess {
		HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		method, ok := dandlerObjMap[bbReq.Bin.FunId]
		if ok {
			log.Info("HandleDataBin %d", bbReq.Bin.FunId)
			responseBytes, code, errMsg := method(bbReq)
			if code == webcommon.ResponseCodeSuccess {
				HandleSuccessResponse(writer, bbReq, responseBytes)
			} else {
				HandleErrorResponse(writer, bbReq, code, errMsg)
			}
		} else {
			HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeInvalidFunId, "Invalid FunId")
		}
	}
}
