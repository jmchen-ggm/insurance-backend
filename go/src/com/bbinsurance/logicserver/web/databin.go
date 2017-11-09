package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"net/http"
)

func HandleDataBin(writer http.ResponseWriter, request *http.Request) {
	log.Info("New Request: %s %s", request.URL, request.Method)
	bbReq, code, msg := HandleRequest(request)
	if code != protocol.ResponseCodeSuccess {
		HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		if bbReq.Bin.FunId == protocol.FuncListArticle {
			log.Info("HandleDataBin ListArticle")
			responseBytes := GetListArticle(bbReq)
			HandleSuccessResponse(writer, bbReq, responseBytes)
		} else {
			HandleSuccessResponse(writer, bbReq, nil)
		}
	}
}
