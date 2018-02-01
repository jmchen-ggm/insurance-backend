package web

import (
	"com/bbinsurance/kvserver/database"
	"com/bbinsurance/webcommon"
	"encoding/json"
)

func FuncKvCreateHttpFun(bbReq webcommon.BBReq) ([]byte, int, string) {
	var httpFunRequest webcommon.KvHttpFunRequest
	json.Unmarshal(bbReq.Body, &httpFunRequest)
	httpFun := database.InsertHttpFun(httpFunRequest.HttpFun)
	if httpFun.Id != -1 {
		var response webcommon.KvHttpFunResponse
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	} else {
		return nil, webcommon.ResponseCodeServerError, ""
	}
}
