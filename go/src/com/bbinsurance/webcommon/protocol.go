package webcommon

import (
	"encoding/json"
)

type BBBin struct {
	FunId     int
	URI       string
	SessionId string
	Timestamp int64
}

type BBReqHeader struct {
	Uin      int64
	Username string
	Token    string
}

type BBRespHeader struct {
	Username     string
	ResponseCode int
	ErrMsg       string
}

type BBResp struct {
	Bin    BBBin
	Header BBRespHeader
	Body   json.RawMessage
}

type BBReq struct {
	Bin    BBBin
	Header BBReqHeader
	Body   json.RawMessage
}
