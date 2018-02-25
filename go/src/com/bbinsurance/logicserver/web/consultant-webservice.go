package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/logicserver/service"
	"com/bbinsurance/webcommon"
	"encoding/json"
)

func FunListConsultant(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listConsultantRequest protocol.BBListConsultantRequest
	json.Unmarshal(bbReq.Body, &listConsultantRequest)
	consultantList := service.GetListConsultant(listConsultantRequest.StartIndex, listConsultantRequest.PageSize)
	log.Info("req %d %d %d", listConsultantRequest.StartIndex, listConsultantRequest.PageSize, len(consultantList))
	var response protocol.BBListConsultantResponse
	response.ConsultantList = consultantList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}
