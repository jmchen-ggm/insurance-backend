package web

import (
	"com/bbinsurance/webcommon"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleRequest(request *http.Request) (webcommon.BBReq, int, string) {
	var bbReq webcommon.BBReq
	if request.Method != "POST" {
		return bbReq, webcommon.ResponseCodeRequestInvalid, "Please Use Http Post"
	} else {
		body, err := ioutil.ReadAll(request.Body)
		if err == nil {
			err = json.Unmarshal(body, &bbReq)
			if err != nil {
				return bbReq, webcommon.ResponseCodeRequestInvalid, "Decode Request Json Err"
			}
		} else {
			return bbReq, webcommon.ResponseCodeRequestInvalid, "Decode Request Json Err"
		}
	}
	return bbReq, webcommon.ResponseCodeSuccess, ""
}

func HandleSuccessResponse(writer http.ResponseWriter, request webcommon.BBReq, body []byte) {
	var bbResp webcommon.BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	if body != nil {
		json.Unmarshal(body, &bbResp.Body)
	}
	responseJsonStr, _ := json.Marshal(bbResp)
	writer.Header().Set("content-type", "application/json")
	fmt.Fprintf(writer, string(responseJsonStr))
}

func HandleErrorResponse(writer http.ResponseWriter, request webcommon.BBReq, errorCode int, errMsg string) {
	var bbResp webcommon.BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Header.ResponseCode = errorCode
	bbResp.Header.ErrMsg = errMsg
	bbResp.Body = *new(json.RawMessage)
	responseJsonStr, _ := json.Marshal(bbResp)
	writer.Header().Set("content-type", "application/json")
	fmt.Fprintf(writer, string(responseJsonStr))
}
