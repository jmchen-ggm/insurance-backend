package webcommon

import (
	"com/bbinsurance/log"
	"com/bbinsurance/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleRequest(request *http.Request) (BBReq, int, string) {
	var bbReq BBReq
	if request.Method != "POST" {
		return bbReq, ResponseCodeRequestInvalid, "Please Use Http Post"
	} else {
		body, err := ioutil.ReadAll(request.Body)
		if err == nil {
			err = json.Unmarshal(body, &bbReq)
			if err != nil {
				log.Error("receive body %s", util.BytesToString(body))
				return bbReq, ResponseCodeRequestInvalid, "Decode Request Json Err"
			}
		} else {
			return bbReq, ResponseCodeRequestInvalid, "Decode Request Json Err"
		}
	}
	log.Info("HandleRequest %s", bbReq.Body)
	return bbReq, ResponseCodeSuccess, ""
}

func HandleSuccessResponse(writer http.ResponseWriter, request BBReq, body []byte) {
	var bbResp BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	if body != nil {
		json.Unmarshal(body, &bbResp.Body)
	}
	responseJsonBytes, _ := json.Marshal(bbResp)
	responseJsonStr := string(responseJsonBytes)
	log.Info("HandleSuccessResponse responseJsonStr: %s", responseJsonStr)
	writer.Header().Set("content-type", "application/json")
	writer.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With, Token")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Max-Age", "86400")
	fmt.Fprintf(writer, string(responseJsonStr))
}

func HandleErrorResponse(writer http.ResponseWriter, request BBReq, errorCode int, errMsg string) {
	var bbResp BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Header.ResponseCode = errorCode
	bbResp.Header.ErrMsg = errMsg
	bbResp.Body = *new(json.RawMessage)
	responseJsonStr, _ := json.Marshal(bbResp)
	log.Info("HandleErrorResponse code: %d errMsg: %s", errorCode, errMsg)
	writer.Header().Set("content-type", "application/json")
	writer.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With, Token")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Max-Age", "86400")
	fmt.Fprintf(writer, string(responseJsonStr))
}

func GenerateImgFileServerUrl(thumbUrl string) string {
	return FileServer + thumbUrl
}
