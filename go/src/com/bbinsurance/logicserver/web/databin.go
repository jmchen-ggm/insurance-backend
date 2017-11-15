package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"net/http"
)

type HandleMethod func(protocol.BBReq) ([]byte, int, string)

var dandlerObjMap map[int]HandleMethod

func FunInitDataBin() {
	dandlerObjMap = make(map[int]HandleMethod)
	dandlerObjMap[protocol.FuncListArticle] = FunGetListArticle
	dandlerObjMap[protocol.FuncListCompany] = FunGetListCompany
	dandlerObjMap[protocol.FuncListInsurance] = FunGetListInsurance
	dandlerObjMap[protocol.FuncListComment] = FunGetListComment
	dandlerObjMap[protocol.FuncCreateComment] = FunCreateComment
	dandlerObjMap[protocol.FuncViewComment] = FunViewComment
}

func FunHandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := HandleRequest(request)
	if code != protocol.ResponseCodeSuccess {
		HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		method, ok := dandlerObjMap[bbReq.Bin.FunId]
		if ok {
			log.Info("HandleDataBin %d", bbReq.Bin.FunId)
			responseBytes, code, errMsg := method(bbReq)
			if code == protocol.ResponseCodeSuccess {
				HandleSuccessResponse(writer, bbReq, responseBytes)
			} else {
				HandleErrorResponse(writer, bbReq, code, errMsg)
			}
		} else {
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeInvalidFunId, "Invalid FunId")
		}
	}
}
