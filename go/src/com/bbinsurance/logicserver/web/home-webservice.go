package web

import (
	"com/bbinsurance/logicserver/service"
	"com/bbinsurance/webcommon"
	"encoding/json"
)

func FunGetHomeData(bbReq webcommon.BBReq) ([]byte, int, string) {
	response := service.GetHomeData()
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}
