package web

import (
	"com/bbinsurance/logicserver/protocol"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleRequest(request *http.Request) (protocol.BBReq, int, string) {
	var bbReq protocol.BBReq
	if request.Method != "POST" {
		return bbReq, protocol.ResponseCodeRequestInvalid, "Please Use Http Post"
	} else {
		body, err := ioutil.ReadAll(request.Body)
		if err == nil {
			err = json.Unmarshal(body, &bbReq)
			if err != nil {
				return bbReq, protocol.ResponseCodeRequestInvalid, "Decode Request Json Err"
			}
		} else {
			return bbReq, protocol.ResponseCodeRequestInvalid, "Decode Request Json Err"
		}
	}
	return bbReq, protocol.ResponseCodeSuccess, ""
}

func HandleSuccessResponse(writer http.ResponseWriter, request protocol.BBReq, body json.RawMessage) {
	var bbResp protocol.BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Body = body
	responseJsonStr, _ := json.Marshal(bbResp)
	writer.Header().Set("content-type", "application/json")
	fmt.Fprintf(writer, string(responseJsonStr))
}

func HandleErrorResponse(writer http.ResponseWriter, request protocol.BBReq, errorCode int, errMsg string) {
	var bbResp protocol.BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Header.ResponseCode = errorCode
	bbResp.Header.ErrMsg = errMsg
	bbResp.Body = *new(json.RawMessage)
	responseJsonStr, _ := json.Marshal(bbResp)
	writer.Header().Set("content-type", "application/json")
	fmt.Fprintf(writer, string(responseJsonStr))
}
