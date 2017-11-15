package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"net/http"
)

type HandleMethod func(protocol.BBReq) []byte

var HandlerObjMap map[int]HandleMethod

func InitDataBin() {
	HandlerObjMap[protocol.FuncListArticle] = GetListArticle
	HandlerObjMap[protocol.FuncListCompany] = GetListCompany
	HandlerObjMap[protocol.FuncListInsurance] = GetListInsurance
	HandlerObjMap[protocol.FuncListComment] = GetListComment
	HandlerObjMap[protocol.FuncCreateComment] = FuncCreateComment
	HandlerObjMap[protocol.FuncViewComment] = FuncViewComment
}

func HandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := HandleRequest(request)
	if code != protocol.ResponseCodeSuccess {
		HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		method, ok := HandlerObjMap[bReq.Bin.FunId]
		if ok {
			log.Info("HandleDataBin %d", bReq.Bin.FunId)
			responseBytes := method(bbReq)
			HandleSuccessResponse(writer, bbReq, responseBytes)
		} else {
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeInvalidFunId, "Invalid FunId")
		}
	}
}
