package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleRequest(request *http.Request) BBReq {
	var bbReq BBReq
	body, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(body, &bbReq)
	return bbReq
}

func HandleSuccessResponse(writer http.ResponseWriter, request BBReq, body json.RawMessage) {
	var bbResp BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Body = body
	responseJsonStr, _ := json.Marshal(bbResp)
	writer.Header().Set("content-type", "application/json")
	fmt.Fprintf(writer, string(responseJsonStr))
}

func HandleErrorResponse(writer http.ResponseWriter, request BBReq, errorCode int, errMsg string) {
	var bbResp BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Header.ResponseCode = errorCode
	bbResp.Header.ErrMsg = errMsg
	responseJsonStr, _ := json.Marshal(bbResp)
	writer.Header().Set("content-type", "application/json")
	fmt.Fprintf(writer, string(responseJsonStr))
}
