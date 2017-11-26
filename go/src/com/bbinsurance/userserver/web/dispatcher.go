package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/userserver/protocol"
	"com/bbinsurance/webcommon"
	"net/http"
)

type HandleMethod func(protocol.BBReq) ([]byte, int, string)

var dandlerObjMap map[int]HandleMethod

func FunInitDataBin() {
	dandlerObjMap = make(map[int]HandleMethod)
}

func FunHandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := webcommon.HandleRequest(request)
	if code != protocol.ResponseCodeSuccess {
		webcommon.HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		method, ok := dandlerObjMap[bbReq.Bin.FunId]
		if ok {
			log.Info("HandleDataBin %d", bbReq.Bin.FunId)
			responseBytes, code, errMsg := method(bbReq)
			if code == protocol.ResponseCodeSuccess {
				webcommon.HandleSuccessResponse(writer, bbReq, responseBytes)
			} else {
				webcommon.HandleErrorResponse(writer, bbReq, code, errMsg)
			}
		} else {
			webcommon.HandleErrorResponse(writer, bbReq, protocol.ResponseCodeInvalidFunId, "Invalid FunId")
		}
	}
}
