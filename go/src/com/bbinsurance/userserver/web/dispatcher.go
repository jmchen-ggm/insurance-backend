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
	handlerObjMap[webcommon.FuncLogin] = FunLogin
	handlerObjMap[webcommon.FuncGetUser] = FunGetUser
	handlerObjMap[webcommon.FuncBatchGetUser] = FunBatchGetUser
	handlerObjMap[webcommon.FuncGetAllUser] = FunGetAllUser
}

func FunHandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := webcommon.HandleRequest(request)
	if code != webcommon.ResponseCodeSuccess {
		webcommon.HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		method, ok := handlerObjMap[bbReq.Bin.FunId]
		if ok {
			responseBytes, code, errMsg := method(bbReq)
			log.Info("HandleDataBin : %d code : %d errMsg : %s", bbReq.Bin.FunId, code, errMsg)
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
