package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"encoding/json"
	"net/http"
)

func HandleDataBin(writer http.ResponseWriter, request *http.Request) {
	bbReq, code, msg := HandleRequest(request)
	if code != protocol.ResponseCodeSuccess {
		HandleErrorResponse(writer, bbReq, code, msg)
	} else {
		if bbReq.Bin.FunId == protocol.FuncListArticle {
			log.Info("HandleDataBin ListArticle")
			responseRawMessage := GetListArticle(bbReq)
			HandleSuccessResponse(writer, bbReq, responseRawMessage)
		} else {
			var responseRawMessage json.RawMessage
			HandleSuccessResponse(writer, bbReq, responseRawMessage)
		}
	}
}
