package webcommon

import (
	"bytes"
	"com/bbinsurance/log"
	"com/bbinsurance/time"
	"com/bbinsurance/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var requestHandleMap = make(map[string]int64)

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
	log.Info("HandleRequest %s", util.ObjToString(bbReq))
	inRequestMap(bbReq)
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
	fmt.Fprintf(writer, responseJsonStr)
	outRequestMap(request, len(responseJsonStr))
}

func HandleErrorResponse(writer http.ResponseWriter, request BBReq, errorCode int, errMsg string) {
	var bbResp BBResp
	bbResp.Bin = request.Bin
	bbResp.Header.Username = request.Header.Username
	bbResp.Header.ResponseCode = errorCode
	bbResp.Header.ErrMsg = errMsg
	bbResp.Body = *new(json.RawMessage)
	responseJsonBytes, _ := json.Marshal(bbResp)
	log.Info("HandleErrorResponse code: %d errMsg: %s", errorCode, errMsg)
	writer.Header().Set("content-type", "application/json")
	writer.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With, Token")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Max-Age", "86400")
	fmt.Fprintf(writer, string(responseJsonBytes))
	outRequestMap(request, len(responseJsonBytes))
}

func inRequestMap(bbReq BBReq) {
	key := fmt.Sprintf("%d$%s", bbReq.Bin.FunId, bbReq.Bin.SessionId)
	requestHandleMap[key] = time.GetTimestampInMilli()
}

func outRequestMap(bbReq BBReq, responseSize int) {
	key := fmt.Sprintf("%d$%s", bbReq.Bin.FunId, bbReq.Bin.SessionId)
	requestTime := requestHandleMap[key]
	responseTime := time.GetTimestampInMilli()
	useTime := responseTime - requestTime
	delete(requestHandleMap, key)

	log.Info("Total Use Time key:%s useTime:%d mapSize:%d responseSize:%d", key, useTime, len(requestHandleMap), responseSize)

	if bbReq.Bin.FunId < FuncKv {
		var httpFun HttpFun
		httpFun.FunId = bbReq.Bin.FunId
		httpFun.Timestamp = requestTime
		httpFun.ResponseSize = responseSize
		httpFun.UseTime = useTime
		httpFun.Uin = bbReq.Header.Uin
		kvHttpFun(httpFun)
	}
}

func GenerateImgFileServerUrl(thumbUrl string) string {
	return FileServer + thumbUrl
}

func kvHttpFun(httpFun HttpFun) {
	var bbReq BBReq
	bbReq.Bin.FunId = FuncKvHttpFun
	bbReq.Bin.URI = UriKv
	var httpFunRequest KvHttpFunRequest
	httpFunRequest.HttpFun = httpFun
	body, _ := json.Marshal(httpFunRequest)
	if body != nil {
		json.Unmarshal(body, &bbReq.Body)
	}
	requestData, _ := json.Marshal(bbReq)
	requestBuffer := bytes.NewBuffer([]byte(requestData))
	requestBodyType := "application/json;charset=utf-8"
	http.Post("http://127.0.0.1:8083/data-bin", requestBodyType, requestBuffer)
}
